package router

import "github.com/gin-gonic/gin"

// Initialize sets up the router and starts the server
func Initialize() {
	// Initialize the router
	router := gin.Default()

	// Set up routes
	initializeRoutes(router)

	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
