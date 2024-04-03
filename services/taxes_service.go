package service

import (
	"goback/database"
	"goback/models"
)

type TaxesService struct{}

func (s *TaxesService) CreateTaxes(taxe models.Taxes) error {
	return database.DB.Create(&taxe).Error
}

func (s *TaxesService) GetTaxes() ([]models.Taxes, error) {
	var taxe []models.Taxes
	if err := database.DB.Find(&taxe).Error; err != nil {
		return nil, err
	}
	return taxe, nil
}

func (s *TaxesService) CreateMultipleTaxes(taxes []models.Taxes) error {

	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, tax := range taxes {
		if err := tx.Create(&tax).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
