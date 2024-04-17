package routes

import (
	"goback/controllers"
	service "goback/services"

	"github.com/gin-gonic/gin"
)

func ProductRouter(router *gin.Engine, productService *service.ProductService) {
	productController := controllers.NewProductController(productService)

	routes := router.Group("/api/products")

	routes.POST("", productController.AddProduct)

}
