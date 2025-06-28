package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySqlGormDB() (*gorm.DB, error) {
	host := getEnv("DB_HOST", "127.0.0.1")
	port := getEnv("DB_PORT", "3306")
	user := getEnv("DB_USER", "root")
	password := getEnv("DB_PASSWORD", "secret")
	dbname := getEnv("DB_NAME", "registration")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	database, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening DB: %s", err)
	}

	fmt.Println("Connected to MySQL Gorm DB")

	return database, nil
}
