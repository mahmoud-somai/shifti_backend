package controllers

import (
	"goback/models"
	service "goback/services"

	"github.com/gin-gonic/gin"
)

type OrderNoteController struct {
	OrderNoteService *service.OrderNoteService
}

func NewOrderNoteController(service *service.OrderNoteService) *OrderNoteController {
	return &OrderNoteController{OrderNoteService: service}
}

func (controller *OrderNoteController) OrderNoteGet(c *gin.Context) {
	notes, err := controller.OrderNoteService.GetOrderNotes()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve order notes"})
		return
	}
	c.JSON(200, notes)
}

// func (controller *OrderNoteController) OrderNoteCreate(c *gin.Context) {
// 	var requestBody models.OrderNote
// 	if err := c.BindJSON(&requestBody); err != nil {
// 		c.JSON(400, gin.H{"error": "Invalid request body"})
// 		return
// 	}
// 	err := controller.OrderNoteService.CreateOrderNote(requestBody)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Failed to create order note"})
// 		return
// 	}
// 	c.JSON(200, requestBody)
// }

func (controller *OrderNoteController) OrderNoteCreate(c *gin.Context) {
	var order_notes []models.OrderNote
	if err := c.BindJSON(&order_notes); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	for _, order_note := range order_notes {
		if err := controller.OrderNoteService.CreateOrderNote(order_note); err != nil {
			c.JSON(500, gin.H{"error": "Failed to add order note"})
			return
		}
	}

	c.JSON(200, gin.H{"message": "order notes insered successfully"})
}
