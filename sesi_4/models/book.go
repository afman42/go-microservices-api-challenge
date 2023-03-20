package models

import "time"

type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	NameBook  string    `gorm:"not null;type:varchar(191)" json:"name_book"`
	Author    string    `gorm:"not null;type:varchar(191)" json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}
