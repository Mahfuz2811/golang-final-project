package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySqlGormDB() (*gorm.DB, error) {
	connectionString := "root:secret@tcp(127.0.0.1:3306)/registration?parseTime=true"

	database, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening DB: %s", err)
	}

	fmt.Println("Connected to MySQL Gorm DB")

	return database, nil
}
