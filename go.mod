module github.com/marcusburghardt/openscap-prototype

go 1.22.8

require (
	github.com/marcusburghardt/comply-prototype v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.67.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
)

replace github.com/marcusburghardt/comply-prototype => ../comply-prototype
