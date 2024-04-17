package main

import (
	"goback/database"
	"goback/models"
)

func init() {
	database.ConnectToDB()
}

func main() {

	database.DB.AutoMigrate(&models.OrderNote{})
	database.DB.AutoMigrate(&models.Taxes{})
	database.DB.AutoMigrate(&models.Dimension{})

	database.DB.AutoMigrate(&models.Product{})
	database.DB.AutoMigrate(&models.Customer{})
	database.DB.AutoMigrate(&models.Order{})

	database.DB.AutoMigrate(&models.Tag{})
	database.DB.AutoMigrate(&models.Image{})
	database.DB.AutoMigrate(&models.Attribute{})

	database.DB.AutoMigrate(&models.Cat{})
	database.DB.AutoMigrate(&models.Category{})

}
