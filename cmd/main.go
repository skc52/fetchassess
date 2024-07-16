package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/skc52/FETCH_TAKE_HOME/cmd/api"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Read the port from the environment variable or use a default value
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8081" // Default port if PORT environment variable is not set
	}

	// Create a new server with the specified port
	server := api.NewAPIServer(":" + portStr)

	// Run the server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
