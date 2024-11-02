package server

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/marcusburghardt/comply-prototype/proto"
	"github.com/marcusburghardt/openscap-prototype/config"
	"github.com/marcusburghardt/openscap-prototype/scan"
	"google.golang.org/grpc"
)

type PluginServer struct {
	proto.UnimplementedScanServiceServer
	Config *config.Config
}

func (s *PluginServer) Execute(ctx context.Context, req *proto.ScanRequest) (*proto.ScanResponse, error) {
	action := req.GetAction()

	switch action {
	case "scan":
		rc, err := scan.ScanSystem(s.Config, "cis")
		if err != nil {
			log.Printf("Error executing command: %v", err)
			return &proto.ScanResponse{ReturnCode: rc}, nil
		}
	case "remediate":
		err := fmt.Errorf("action not yet implemented: %s", action)
		return &proto.ScanResponse{ReturnCode: 1}, err
	default:
		err := fmt.Errorf("unknown action: %s", action)
		return &proto.ScanResponse{ReturnCode: 1}, err
	}

	return &proto.ScanResponse{ReturnCode: 0}, nil
}

func isSocketFile(fileInfo os.FileInfo) (bool, error) {
	if fileInfo.Mode()&os.ModeSocket != 1 {
		return false, fmt.Errorf("there is an existing file using the socket path: %s", fileInfo.Name())
	}
	return true, nil
}

func setupListener(socket string) (net.Listener, error) {
	if stat, err := os.Stat(socket); err == nil {
		if _, err := isSocketFile(stat); err != nil {
			return nil, err
		}
		if err := os.Remove(socket); err != nil {
			return nil, fmt.Errorf("failed to remove existing socket: %w", err)
		}
	}
	listener, err := net.Listen("unix", socket)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on socket: %w", err)
	}
	return listener, nil
}

func cleanupSocket(socket string) {
	if stat, err := os.Stat(socket); err == nil {
		if _, err := isSocketFile(stat); err != nil {
			log.Printf("Error removing socket file: %v", err)
		} else if err := os.Remove(socket); err != nil && !errors.Is(err, fs.ErrNotExist) {
			log.Printf("Error removing socket file: %v", err)
		}
	}
}

func createGRPCServer(cfg *config.Config) *grpc.Server {
	grpcServer := grpc.NewServer()
	pluginServer := &PluginServer{Config: cfg}
	proto.RegisterScanServiceServer(grpcServer, pluginServer)
	return grpcServer
}

func runServer(grpcServer *grpc.Server, listener net.Listener, socket string) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a separate goroutine so we can better control the execution
	go func() {
		log.Printf("Server listening on Unix socket %s", socket)
		if err := grpcServer.Serve(listener); err != nil && err != grpc.ErrServerStopped {
			log.Fatalf("Failed to serve gRPC request: %v", err)
		}
	}()

	<-stop
	grpcServer.GracefulStop()
}

func getSocketPath(socket string) (string, error) {
	socketFile, err := config.SanitizeInput(socket)
	if err != nil {
		return "", err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	socketPath := homeDir + "/" + socketFile
	return socketPath, nil
}

func StartServer(cfg *config.Config) error {
	socket, err := getSocketPath(cfg.Server.Socket)
	if err != nil {
		return err
	}

	listener, err := setupListener(socket)
	if err != nil {
		return fmt.Errorf("failed to setup listener: %v", err)
	}
	defer cleanupSocket(socket)

	grpcServer := createGRPCServer(cfg)
	runServer(grpcServer, listener, socket)
	return nil
}
