package routes

import (
	"goback/controllers"
	service "goback/services"

	"github.com/gin-gonic/gin"
)

func OrderNoteRouter(router *gin.Engine, orderNoteService *service.OrderNoteService) {
	orderNoteController := controllers.NewOrderNoteController(orderNoteService)

	routes := router.Group("api/ordersnote")

	routes.GET("", orderNoteController.OrderNoteGet)
	routes.POST("", orderNoteController.OrderNoteCreate)
}
