package controllers

import (
	configs "goback/database"
	"goback/models"
	"time"

	"github.com/gin-gonic/gin"
)

func OrderNoteGet(c *gin.Context) {
	var orders_notes []models.OrderNote
	configs.DB.Find(&orders_notes)
	c.JSON(200, &orders_notes)
	return
}

type OrderNoteRequestBody struct {
	ForeignId uint      `json:"id_foreign"`
	Date      time.Time `json:"date"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
}

// func OrderNoteCreate(c *gin.Context) {

// 	body := OrderNoteRequestBody{}

// 	c.BindJSON(&body)

// 	orderNote := &models.OrderNote{ForeignId: body.ForeignId, Date: body.Date, Author: body.Author, Content: body.Content}

// 	result := configs.DB.Create(&orderNote)

// 	if result.Error != nil {
// 		c.JSON(500, gin.H{"Error": "Failed to insert"})
// 		return
// 	}

// 	c.JSON(200, &orderNote)
// }

func OrderNoteCreate(c *gin.Context) {
	var requestBody []OrderNoteRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		// If there's an error, check if it's due to single JSON object
		var singleBody OrderNoteRequestBody
		if err := c.BindJSON(&singleBody); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		// Handle single object case
		requestBody = append(requestBody, singleBody)
	}

	var orderNotes []models.OrderNote
	for _, body := range requestBody {
		orderNote := &models.OrderNote{
			ForeignId: body.ForeignId,
			Date:      body.Date,
			Author:    body.Author,
			Content:   body.Content,
		}
		orderNotes = append(orderNotes, *orderNote)
	}

	result := configs.DB.Create(&orderNotes)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to insert"})
		return
	}

	c.JSON(200, &orderNotes)
}
