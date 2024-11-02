package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/marcusburghardt/openscap-prototype/config"
	"github.com/marcusburghardt/openscap-prototype/server"
)

func parseFlags() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "./oscap-config.yml", "Path to config file")
	flag.Parse()

	configFile, err := config.SanitizeAndValidatePath(configPath, false)
	if err != nil {
		return "", err
	}

	return configFile, nil
}

func initializeConfig() (*config.Config, error) {
	configFile, err := parseFlags()
	if err != nil {
		return nil, fmt.Errorf("error parsing flags: %w", err)
	}

	config, err := config.ReadConfig(configFile)
	if err != nil {
		return nil, fmt.Errorf("error reading config from %s: %w", configFile, err)
	}

	return config, nil
}

func main() {
	config, err := initializeConfig()
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	err = server.StartServer(config)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
