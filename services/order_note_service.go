package service

import (
	"goback/database"
	"goback/models"
)

// OrderNoteService handles operations related to order notes.
type OrderNoteService struct{}

// CreateOrderNote creates a new order note.
func (s *OrderNoteService) CreateOrderNote(note models.OrderNote) error {
	return database.DB.Create(&note).Error
}

// GetOrderNotes retrieves all order notes from the database.
func (s *OrderNoteService) GetOrderNotes() ([]models.OrderNote, error) {
	var notes []models.OrderNote
	if err := database.DB.Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

// Other service functions can be defined similarly.
