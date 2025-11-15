// Package server provides HTTP server functionality for the dev echo server.
// It includes middleware support for logging, security headers, and error handling.
package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/islamMaruf/dev-echo-server/middleware"
	"github.com/islamMaruf/dev-echo-server/routes"
)

// Server represents an HTTP server instance with configured routes and middleware.
type Server struct {
	port   string
	router *http.ServeMux
}

// NewServer creates and initializes a new HTTP server instance.
// It configures the server with the specified port and sets up all routes.
//
// Parameters:
//   - port: The port number to listen on (e.g., "3000" or "8080")
//
// Returns:
//   - *Server: A configured server instance ready to start
//
// Example:
//
//	srv := server.NewServer("8080")
//	if err := srv.Start(); err != nil {
//	    log.Fatal(err)
//	}
func NewServer(port string) *Server {
	s := &Server{
		port:   port,
		router: http.NewServeMux(),
	}
	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	// Register routes
	routes.RegisterRoutes(s.router)
}

// Start starts the HTTP server with all configured middleware.
// It sets up logging to files, applies security headers, and error handling.
// The server will listen on the port specified during initialization.
//
// Returns:
//   - error: Any error encountered while starting the server, or nil on success
//
// The server includes:
//   - Request logging with daily log file rotation
//   - Security headers (Helmet.js-like protection)
//   - Error handling with panic recovery
//   - Configurable timeouts (15s read/write, 60s idle)
func (s *Server) Start() error {
	// Setup logging
	logDir := "log"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Printf("Failed to create log directory: %v", err)
	}

	logFile, err := os.OpenFile(
		filepath.Join(logDir, fmt.Sprintf("access-%s.log", time.Now().Format("2006-01-02"))),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// Chain middleware
	handler := middleware.Logger(logFile)(s.router)
	handler = middleware.Security(handler)
	handler = middleware.ErrorHandler(handler)

	// Create and start server
	server := &http.Server{
		Addr:         ":" + s.port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return server.ListenAndServe()
}
