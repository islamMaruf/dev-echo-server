package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"dev-echo-server/middleware"
	"dev-echo-server/routes"
)

type Server struct {
	port   string
	router *http.ServeMux
}

// NewServer creates a new HTTP server
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

// Start starts the HTTP server with middleware
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
