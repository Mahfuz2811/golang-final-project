package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLGormDB() (*gorm.DB, error) {
	dsn := "root:secret@tcp(127.0.0.1:3306)/registration?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	fmt.Println("Connected to MySQL with GORM")
	return db, nil
}
