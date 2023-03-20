package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDB() *gorm.DB {
	const (
		HOST         = "localhost"
		USER         = "postgres"
		PASSWORD     = "afif123"
		DATABASEPORT = "5432"
		DATABASENAME = "showcase-gorm"
	)

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASSWORD, DATABASENAME, DATABASEPORT)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	return db
}
