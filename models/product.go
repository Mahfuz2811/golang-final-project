package models

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"type:text"`
	Price       float64
	UserEmail   string `gorm:"index;not null"` // FK by email for simplicity
}
