package routes

import (
	"goback/controllers"

	"github.com/gin-gonic/gin"
)

func OrderNoteRouter(router *gin.Engine) {

	routes := router.Group("api/ordersnote")

	routes.GET("", controllers.OrderNoteGet)
	routes.POST("", controllers.OrderNoteCreate)

}
