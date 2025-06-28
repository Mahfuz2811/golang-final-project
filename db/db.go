package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func NewMySqlDB() (*sql.DB, error) {
	host := getEnv("DB_HOST", "127.0.0.1")
	port := getEnv("DB_PORT", "3306")
	user := getEnv("DB_USER", "root")
	password := getEnv("DB_PASSWORD", "secret")
	dbname := getEnv("DB_NAME", "registration")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	database, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error opening DB: %s", err)
	}

	if err := database.Ping(); err != nil {
		return nil, fmt.Errorf("error opening DB: %s", err)
	}

	DB = database

	fmt.Println("Connected to MySQL DB")

	return database, nil
}
