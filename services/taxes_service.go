package service

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"goback/database"
	"goback/models"
)

var ErrInvalidScanType = errors.New("invalid scan type")

type TaxesService struct{}

type StringSlice []string

func (s *StringSlice) Scan(value interface{}) error {
	if value == nil {
		*s = nil
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	default:
		return ErrInvalidScanType
	}

	return json.Unmarshal([]byte(str), s)
}

func (s StringSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *TaxesService) CreateTaxes(taxe models.Taxes) error {
	return database.DB.Create(&taxe).Error
}

func (s *TaxesService) GetTaxes() ([]models.Taxes, error) {
	var taxes []models.Taxes
	if err := database.DB.Find(&taxes).Error; err != nil {
		return nil, err
	}
	return taxes, nil
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
