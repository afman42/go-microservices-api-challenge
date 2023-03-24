package main

import (
	"sesi_4_project/database"
	"sesi_4_project/models"
	"sesi_4_project/routers"
)

func main() {
	db := database.LoadDB()
	db.AutoMigrate(&models.Book{})

	r := routers.LoadRoutes(db)
	r.Run(":8080")
}
