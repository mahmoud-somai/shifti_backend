package routes

import (
	"goback/controllers"
	service "goback/services"

	"github.com/gin-gonic/gin"
)

func OrderNoteRouter(router *gin.Engine, orderNoteService *service.OrderNoteService) {
	orderNoteController := controllers.NewOrderNoteController(orderNoteService) // Create an instance of the controller

	routes := router.Group("api/ordersnote")

	// Use controller methods as route handlers
	routes.GET("", orderNoteController.OrderNoteGet)
	routes.POST("", orderNoteController.OrderNoteCreate)
}
