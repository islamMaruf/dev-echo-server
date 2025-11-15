# Dev Echo Server v1.0.0 - Initial Release 🎉

A lightweight, developer-friendly HTTP echo server for testing webhooks, debugging API integrations, and development workflows. Built with Go, featuring Docker support, comprehensive logging, and security middleware.

## 🚀 Highlights

- **Easy to Use**: Run with a single command using the interactive `run.sh` script
- **Docker Ready**: Full Docker and Docker Compose support with multi-stage builds
- **Production Features**: Security headers, request logging, error handling, and panic recovery
- **Flexible**: Use as a standalone server or import as a Go library
- **Well Documented**: Comprehensive README with examples and quick-start guides

## ✨ Features

### Core Functionality
- ✅ HTTP echo server that mirrors request data back
- ✅ Catch-all route for any endpoint
- ✅ Request body logging and echoing
- ✅ JSON request/response handling
- ✅ Automatic request ID generation for tracking

### Middleware & Security
- ✅ Security headers (Helmet.js-like for Go)
  - X-Content-Type-Options: nosniff
  - X-Frame-Options: DENY
  - X-XSS-Protection
  - Strict-Transport-Security
  - Content-Security-Policy
- ✅ Request logging with file rotation
- ✅ Error handling with panic recovery
- ✅ Response time tracking

### Configuration & Deployment
- ✅ Environment-based configuration (.env support)
- ✅ Configurable port (default: 3000)
- ✅ Development/Production modes
- ✅ Docker support with optimized builds
- ✅ Docker Compose for easy orchestration
- ✅ Health check endpoints

### Developer Experience
- ✅ Interactive run script with automatic Docker fallback
- ✅ Makefile with common tasks
- ✅ Support for both `docker compose` and `docker-compose`
- ✅ Automatic .env file creation from example
- ✅ Clean, modular project structure
- ✅ Comprehensive documentation

## 📦 Installation

### Option 1: Install as a CLI Tool (Recommended for quick usage)

```bash
go install github.com/islamMaruf/dev-echo-server@v1.0.0
```

Then run:
```bash
dev-echo-server
```

### Option 2: Use as a Go Library

```bash
go get github.com/islamMaruf/dev-echo-server@v1.0.0
```

Example usage:
```go
package main

import (
    "log"
    "github.com/islamMaruf/dev-echo-server/server"
)

func main() {
    srv := server.NewServer("8080")
    log.Println("Starting echo server on port 8080")
    if err := srv.Start(); err != nil {
        log.Fatal(err)
    }
}
```

### Option 3: Clone and Run

```bash
git clone https://github.com/islamMaruf/dev-echo-server.git
cd dev-echo-server
./run.sh
```

## 🎯 Quick Start

### Using the Interactive Script

```bash
./run.sh
```

Choose from:
1. Docker (Recommended - No Go installation needed)
2. Native Go (Requires Go installed)
3. Build Docker image only
4. Stop and remove containers

### Using Docker Compose

```bash
docker-compose up --build -d
```

### Using Makefile

```bash
make help           # Show all commands
make quick          # Interactive script
make docker-up      # Start with Docker
make dev            # Run in development mode
```

## 🔧 Configuration

Create a `.env` file (automatically created from `.env.example`):

```bash
PORT=3000
NODE_ENV=development
```

## 📖 API Endpoints

- `GET /` - Returns welcome message
- `* /*` - Catch-all route that echoes request body back

### Example Request

```bash
curl -X POST http://localhost:3000/webhook \
  -H "Content-Type: application/json" \
  -d '{"event":"test","data":"hello"}'
```

### Example Response

```json
{
  "response": {
    "data": {
      "event": "test",
      "data": "hello"
    },
    "message": "Redirect Data"
  }
}
```

## 🐳 Docker Usage

### Docker Compose (Easiest)

```bash
# Start
docker-compose up -d

# View logs
docker-compose logs -f

# Stop
docker-compose down
```

### Dockerfile Only

```bash
# Build
docker build -t dev-echo-server .

# Run
docker run -d -p 3000:3000 \
  -e PORT=3000 \
  -e NODE_ENV=development \
  --name dev-echo-server \
  dev-echo-server
```

## 📁 Project Structure

```
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
├── Dockerfile           # Multi-stage Docker build
├── docker-compose.yml   # Docker Compose configuration
├── run.sh               # Interactive runner script
├── Makefile             # Common development tasks
├── .env.example         # Environment variables template
└── README.md            # Full documentation
```

## 🛠️ Development

```bash
# Run tests
make test

# Run with live reload (requires entr)
find . -name "*.go" | entr -r go run main.go

# Lint code
make lint

# Clean build artifacts
make clean
```

## 📝 Use Cases

Perfect for:
- Testing webhook integrations locally
- Debugging HTTP requests and payloads
- API development and testing
- Learning Go web development
- Quick HTTP endpoint mockups
- CI/CD pipeline testing
- Request/response logging and inspection

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- **Repository**: https://github.com/islamMaruf/dev-echo-server
- **Documentation**: https://pkg.go.dev/github.com/islamMaruf/dev-echo-server
- **Issues**: https://github.com/islamMaruf/dev-echo-server/issues
- **Changelog**: [CHANGELOG.md](CHANGELOG.md)

## 🙏 Acknowledgments

Built with:
- [Go](https://golang.org/) - The Go Programming Language
- [UUID](https://github.com/google/uuid) - Google UUID package
- [godotenv](https://github.com/joho/godotenv) - Environment variable loading

---

**Full Changelog**: https://github.com/islamMaruf/dev-echo-server/commits/v1.0.0

Made with ❤️ for developers by developers
