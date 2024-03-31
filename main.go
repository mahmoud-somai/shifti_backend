package main

import (
	"goback/database"
	"goback/routes"
	service "goback/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize database connection
	database.ConnectToDB()

	// Initialize services
	orderNoteService := &service.OrderNoteService{} // Initialize your service(s)

	// Initialize routes with service instances
	routes.OrderNoteRouter(r, orderNoteService)

	// Define a default route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})

	// Run the server
	r.Run()
}
