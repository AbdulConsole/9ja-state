package main

import (
	"log"
	"9ja-state/handlers"
	"9ja-state/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load data
	if err := handlers.LoadData(); err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Register versioned routes
	routes.RegisterV1Routes(router)

	// Start the server
	log.Println("Server is running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
