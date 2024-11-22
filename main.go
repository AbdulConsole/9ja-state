package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var data Data

func main() {
	// Load data
	data = LoadData("data.json")

	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.GET("/states", getStates)
	router.GET("/states/:state/local-governments", getLocalGovernments)

	// Run server
	router.Run(":8080")
}

func getStates(c *gin.Context) {
	// Extract states for Nigeria
	states := []string{}
	for state := range data.Countries["Nigeria"] {
		states = append(states, state)
	}

	c.JSON(http.StatusOK, gin.H{
		"states": states,
	})
}

func getLocalGovernments(c *gin.Context) {
	state := c.Param("state")
	lgas, exists := data.Countries["Nigeria"][state]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "State not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"state":            state,
		"local_governments": lgas,
	})
}
