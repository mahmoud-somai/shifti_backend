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

	database.ConnectToDB()

	orderNoteService := &service.OrderNoteService{}
	taxesService := &service.TaxesService{}
	productService := &service.ProductService{}

	routes.OrderNoteRouter(r, orderNoteService)
	routes.TaxesRouter(r, taxesService)
	routes.ProductRouter(r, productService)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})

	r.Run()
}
