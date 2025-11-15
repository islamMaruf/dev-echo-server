# Dev Echo Server

[![Go Version](https://img.shields.io/github/go-mod/go-version/islamMaruf/dev-echo-server)](https://github.com/islamMaruf/dev-echo-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/islamMaruf/dev-echo-server)](https://goreportcard.com/report/github.com/islamMaruf/dev-echo-server)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub release](https://img.shields.io/github/release/islamMaruf/dev-echo-server.svg)](https://github.com/islamMaruf/dev-echo-server/releases)

Development HTTP echo server that mirrors request data back. Useful for testing webhooks, API integrations, and debugging HTTP requests during development.

## Installation

### As a Go Package

```bash
go get github.com/islamMaruf/dev-echo-server
```

### As a Standalone Tool

```bash
# Install globally
go install github.com/islamMaruf/dev-echo-server@latest

# Run it
dev-echo-server
```

### Clone and Run

```bash
git clone https://github.com/islamMaruf/dev-echo-server.git
cd dev-echo-server
./run.sh
```

## Features

- ✅ Security headers (similar to Helmet.js)
- ✅ JSON request/response logging with rotation
- ✅ Error handling middleware
- ✅ Environment-based configuration
- ✅ Clean project structure

## Usage

### As a Library

You can import and use the server components in your own Go applications:

```go
package main

import (
    "log"
    "github.com/islamMaruf/dev-echo-server/server"
)

func main() {
    srv := server.NewServer("8080")
    log.Printf("Starting echo server on port 8080")
    if err := srv.Start(); err != nil {
        log.Fatal(err)
    }
}
```

### As a Standalone Application

See the Quick Start section below.

## Getting Started

### Quick Start (One Command) 🚀

The easiest way to run the server:

```bash
./run.sh
```

This interactive script will:

- Check for `.env` file and create it if needed
- Let you choose between Docker or native Go
- Automatically fall back to native Go if Docker is not installed
- Build and start the server automatically
- Provide clear instructions for stopping

Or use the Makefile:

```bash
make quick    # Interactive script
make help     # Show all available commands
```

### Environment Setup

Copy the example environment file and configure as needed:

```bash
cp .env.example .env
```

Edit `.env` to set your preferred values:

- `PORT` - Server port (default: 3000)
- `NODE_ENV` - Environment mode (development/production)

### Docker Setup 🐳

#### Using Docker Compose (Recommended)

```bash
# Build and start
docker-compose up --build -d

# Or using Makefile
make docker-up

# View logs
docker-compose logs -f
# Or: make docker-logs

# Stop
docker-compose down
# Or: make docker-down
```

#### Using Dockerfile directly

```bash
# Build image
docker build -t dev-echo-server .

# Run container
docker run -d \
  -p 3000:3000 \
  -e PORT=3000 \
  -e NODE_ENV=development \
  --name dev-echo-server \
  dev-echo-server

# Stop container
docker stop dev-echo-server
docker rm dev-echo-server
```

### Manual Setup (Without Docker)

#### Install dependencies

```bash
go mod download
```

### Running in development

```bash
NODE_ENV=development go run main.go
```

### Running in production

```bash
NODE_ENV=production go run main.go
```

### Build and run binary

```bash
# Build the application
go build -o dev-echo-server .

# Run the binary
./dev-echo-server

# Run in background
./dev-echo-server &

# Run with custom port
PORT=8080 ./dev-echo-server
```

### Stop the server

```bash
# If running in foreground, press Ctrl+C

# If running in background
pkill -f dev-echo-server

# Or find and kill by PID
ps aux | grep dev-echo-server
kill <PID>
```

Runs on localhost:3000 by default but can be configured using the `PORT` environment variable.

### Running tests

```bash
go test ./...

# With coverage
go test -cover ./...

# Watch mode (requires entr)
find . -name "*.go" | entr -c go test ./...
```

### Linting

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run

# Fix issues
golangci-lint run --fix
```

## Project Structure

```text
dev-echo-server/
├── main.go              # Application entry point
├── server/
│   └── server.go        # Server setup and configuration
├── routes/
│   └── routes.go        # Route handlers
├── middleware/
│   ├── logger.go        # Request logging middleware
│   ├── error.go         # Error handling middleware
│   └── security.go      # Security headers middleware
├── log/                 # Log files directory
├── .env                 # Environment variables (not in git)
├── .env.example         # Example environment file
├── .dockerignore       # Docker ignore rules
├── Dockerfile          # Docker image definition
├── docker-compose.yml  # Docker compose configuration
├── Makefile            # Common tasks and commands
├── run.sh              # One-command runner script
├── .gitignore          # Git ignore rules
├── go.mod              # Go module file
├── go.sum              # Go dependencies checksum
└── README.md           # This file
```

## API Endpoints

- `GET /` - Returns welcome message
- `* /*` - Catch-all route that echoes request body

## Environment Variables

- `PORT` - Server port (default: 3000)
- `NODE_ENV` - Environment mode (development/production)

## Quick Reference

### Common Commands

| Command | Description |
|---------|-------------|
| `./run.sh` | Interactive runner (easiest way) |
| `make help` | Show all make commands |
| `make quick` | Run interactive script |
| `make docker-up` | Start with Docker |
| `make docker-down` | Stop Docker containers |
| `make docker-logs` | View Docker logs |
| `make build` | Build Go binary |
| `make dev` | Run in development mode |
| `make test` | Run tests |
| `make clean` | Clean build artifacts |

### Health Check

```bash
# Check if server is running
curl http://localhost:3000/

# Test echo endpoint
curl -X POST http://localhost:3000/test \
  -H "Content-Type: application/json" \
  -d '{"message":"hello"}'
```

## License

MIT License - See LICENSE file for details
