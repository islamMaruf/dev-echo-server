package main

import (
	"log"
	"os"

	"github.com/islamMaruf/dev-echo-server/server"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	srv := server.NewServer(port)
	log.Printf("Listening on port %s\n", port)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
