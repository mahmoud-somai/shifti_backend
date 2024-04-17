// services/product_service.go
package service

import (
	"goback/database"

	"goback/models"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) AddProduct(product models.Product) error {
	return database.DB.Create(&product).Error
}
