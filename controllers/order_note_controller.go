package controllers

import (
	"goback/models"
	service "goback/services" // Import the service package

	"github.com/gin-gonic/gin"
)

type OrderNoteController struct {
	OrderNoteService *service.OrderNoteService // Inject the service
}

// Constructor function to create an instance of OrderNoteController
func NewOrderNoteController(service *service.OrderNoteService) *OrderNoteController {
	return &OrderNoteController{OrderNoteService: service}
}

// Modify your controller functions to use the service methods
func (controller *OrderNoteController) OrderNoteGet(c *gin.Context) {
	notes, err := controller.OrderNoteService.GetOrderNotes()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve order notes"})
		return
	}
	c.JSON(200, notes)
}

func (controller *OrderNoteController) OrderNoteCreate(c *gin.Context) {
	var requestBody models.OrderNote
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	err := controller.OrderNoteService.CreateOrderNote(requestBody)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create order note"})
		return
	}
	c.JSON(200, requestBody)
}
