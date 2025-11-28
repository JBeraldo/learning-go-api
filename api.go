package main

import (
	"jberaldo/go-api/database"
	"jberaldo/go-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	routes.RegisterRoutes(r, database.DB)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
