package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDB() *gorm.DB {
	const (
		host         = "localhost"
		user         = "postgres"
		password     = "afif123"
		databasePort = "5432"
		databaseName = "show-case-gorm1"
	)

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, databaseName, databasePort)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	return db
}
