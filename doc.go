// Package main provides a development HTTP echo server for testing and debugging.
//
// Dev Echo Server is a lightweight HTTP server that mirrors request data back to clients.
// It's designed for testing webhooks, debugging API integrations, and development workflows.
//
// # Features
//
//   - HTTP echo server that mirrors request data
//   - Security headers middleware (Helmet.js-like)
//   - Request logging with file rotation
//   - Error handling with panic recovery
//   - Docker support with multi-stage builds
//   - Environment-based configuration
//
// # Quick Start
//
// Install as a CLI tool:
//
//	go install github.com/islamMaruf/dev-echo-server@latest
//	dev-echo-server
//
// Use as a library:
//
//	package main
//
//	import (
//	    "log"
//	    "github.com/islamMaruf/dev-echo-server/server"
//	)
//
//	func main() {
//	    srv := server.NewServer("8080")
//	    log.Println("Starting echo server on port 8080")
//	    if err := srv.Start(); err != nil {
//	        log.Fatal(err)
//	    }
//	}
//
// # Configuration
//
// Configure via environment variables:
//
//   - PORT: Server port (default: 3000)
//   - NODE_ENV: Environment mode (development/production)
//
// # API Endpoints
//
//   - GET / - Returns a welcome message
//   - * /* - Catch-all route that echoes request body
//
// # Example Request
//
//	curl -X POST http://localhost:3000/webhook \
//	  -H "Content-Type: application/json" \
//	  -d '{"event":"test","data":"hello"}'
//
// # Example Response
//
//	{
//	  "response": {
//	    "data": {
//	      "event": "test",
//	      "data": "hello"
//	    },
//	    "message": "Redirect Data"
//	  }
//	}
//
// For more information, visit: https://github.com/islamMaruf/dev-echo-server
package main
