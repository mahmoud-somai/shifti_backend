// // controllers/product_controller.go
// package controllers

// import (
// 	"goback/models"
// 	service "goback/services"

// 	"github.com/gin-gonic/gin"
// )

// type ProductController struct {
// 	ProductService *service.ProductService
// }

// func NewProductController(productService *service.ProductService) *ProductController {
// 	return &ProductController{ProductService: productService}
// }

// func (controller *ProductController) AddProduct(c *gin.Context) {
// 	var product models.Product
// 	if err := c.BindJSON(&product); err != nil {
// 		c.JSON(400, gin.H{"error": "Invalid request body"})
// 		return
// 	}

// 	if err := controller.ProductService.AddProduct(product); err != nil {
// 		c.JSON(500, gin.H{"error": "Failed to add product"})
// 		return
// 	}

// 	c.JSON(200, gin.H{"message": "Product added successfully"})
// }

package controllers

import (
	"goback/models"
	service "goback/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (controller *ProductController) AddProduct(c *gin.Context) {
	var products []models.Product
	if err := c.BindJSON(&products); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	for _, product := range products {
		if err := controller.ProductService.AddProduct(product); err != nil {
			c.JSON(500, gin.H{"error": "Failed to add product"})
			return
		}
	}

	c.JSON(200, gin.H{"message": "Products added successfully"})
}
