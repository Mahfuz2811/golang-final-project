package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model         // Includes ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	Price       float64
	UserEmail   string `gorm:"index;not null"` // FK by email for simplicity
}
