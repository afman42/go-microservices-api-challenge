package models

import "time"

type Book struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	NameBook  string    `gorm:"not null;type:varchar(191)" json:"name_book"`
	Author    string    `gorm:"not null;type:varchar(191)" json:"author"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
