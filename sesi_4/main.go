package main

import (
	"sesi_4_project/models"
	"sesi_4_project/routers"
)

func main() {
	database := models.LoadDB()
	database.AutoMigrate(&models.Book{})

	r := routers.LoadRoutes(database)
	r.Run(":8080")
}
