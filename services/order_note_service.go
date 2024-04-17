package service

import (
	"goback/database"
	"goback/models"
)

type OrderNoteService struct{}

func (s *OrderNoteService) CreateOrderNote(note models.OrderNote) error {
	return database.DB.Create(&note).Error
}

func (s *OrderNoteService) GetOrderNotes() ([]models.OrderNote, error) {
	var notes []models.OrderNote
	if err := database.DB.Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}
