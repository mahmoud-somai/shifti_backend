package main

import (
	configs "goback/database"
	"goback/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.OrderNote{})
	configs.DB.AutoMigrate(&models.Taxes{})
}
