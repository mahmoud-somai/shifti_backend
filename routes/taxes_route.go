package routes

import (
	"goback/controllers"
	service "goback/services"

	"github.com/gin-gonic/gin"
)

func TaxesRouter(router *gin.Engine, TaxesService *service.TaxesService) {
	TaxesController := controllers.NewTaxeController(TaxesService) 

	routes := router.Group("api/taxes")

	
	routes.GET("", TaxesController.TaxeGet)
	routes.POST("", TaxesController.CreateTaxes)
	routes.POST("/multiple", TaxesController.CreateMultipleTaxes)
}
