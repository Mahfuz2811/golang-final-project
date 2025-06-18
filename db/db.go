package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func NewMySqlDB() (*sql.DB, error) {
	connectionString := "root:secret@tcp(127.0.0.1:3306)/registration?parseTime=true"

	database, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error opening DB: %s", err)
	}

	if err := database.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging DB: %s", err)
	}

	DB = database

	fmt.Println("Connected to MySQL DB")

	return database, nil
}
