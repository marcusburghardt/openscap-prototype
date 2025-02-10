# OpenSCAP-Prototype

NOTE: This repository was an experiment that is now finished and therefore is archived.

The outcomes of this experiment were incorporated at https://github.com/complytime/complytime

## Overview

**OpenSCAP-Prototype** is a plugin which extends the **Comply-Prototype** capabilities to use OpenSCAP. The plugin communicates with **Comply-Prototype** via gRPC, providing a standard and consistent communication mechanism while also allowing plugins developers to choose their preferred languages to contribute. This project is structured to allow modular development, ease of packaging, and maintainability.

## Project Structure

```
complytime-prototype/
├── config/             # Package for plugin configuration
│ ├── config_test.go    # Tests for functions in config.go
│ └── config.go         # Main code used to process plugin configuration
├── oscap/              # Package to interact with oscap command
│ └── oscap.go          # Main code used to interact with oscap command
├── scan/               # Package to process system scan instructions
│ └── scan.go           # Main code used to process scan instructions
├── server/             # Package to manage the gRPC server
│ └── server.go         # Main code used to start a gRPC server
├── go.mod              # Go module file
├── main.go             # Main file for Plugin
├── oscap-config.yml    # Plugin config file
└── README.md           # This file
```

## Installation

### Prerequisites

- **Go** version 1.20 or higher
- **Protocol Buffers** compiler (for gRPC)
- **Make** (optional, for using the `Makefile` if included)
- **scap-security-guide** package installed

### Clone the repository

```bash
git clone https://github.com/marcusburghardt/openscap-prototype.git
cd openscap-prototype
```

## Build Instructions
To compile the CLI and plugin:

```bash
go build -o openscap-prototype .
```

## Running
Start the plugin server:

```bash
./openscap-prototype
```

In another terminal, run the **Comply-Prototype** CLI to connect to the server:

```bash
./comply-prototype scan
```

### Testing
Tests are organized within each package. Run tests using:

```bash
go test ./...
```

### Packaging as RPM
To build an RPM package, use the spec file in build/rpm:

```bash
rpmbuild -ba ...
```

## Contributing
Please open an issue or submit a pull request for any contributions or improvements.

## License
MIT License
