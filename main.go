package main

import (
	"net/http"

	configs "goback/database"
	"goback/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Use CORS middleware to allow cross-origin requests
	r.Use(cors.Default())

	// Initialize database connection
	configs.ConnectToDB()

	// Initialize routes
	routes.OrderNoteRouter(r)
	// Define a default route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})

	// Run the server
	r.Run()
}
