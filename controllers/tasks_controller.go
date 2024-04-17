package controllers

import (
	"goback/models"
	service "goback/services"

	"github.com/gin-gonic/gin"
)

type TaxesController struct {
	TaxesService *service.TaxesService
}

func NewTaxeController(service *service.TaxesService) *TaxesController {
	return &TaxesController{TaxesService: service}
}

func (controller *TaxesController) TaxeGet(c *gin.Context) {
	taxes, err := controller.TaxesService.GetTaxes()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve a taxe"})
		return
	}
	c.JSON(200, taxes)
}

func (controller *TaxesController) CreateTaxes(c *gin.Context) {
	var requestBody models.Taxes
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if requestBody.Cities == nil {
		requestBody.Cities = make([]string, 0)
	}
	if requestBody.Postcodes == nil {
		requestBody.Postcodes = make([]string, 0)
	}

	err := controller.TaxesService.CreateTaxes(requestBody)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create a Taxe"})
		return
	}
	c.JSON(200, requestBody)
}

func (controller *TaxesController) CreateMultipleTaxes(c *gin.Context) {
	var requestBody []models.Taxes
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	for i := range requestBody {
		if requestBody[i].Cities == nil {
			requestBody[i].Cities = make([]string, 0)
		}
		if requestBody[i].Postcodes == nil {
			requestBody[i].Postcodes = make([]string, 0)
		}
	}

	err := controller.TaxesService.CreateMultipleTaxes(requestBody)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create taxes"})
		return
	}
	c.JSON(200, requestBody)
}
