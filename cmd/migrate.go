package main

import (
	"final-golang-project/db"
	"final-golang-project/models"
	"flag"
	"fmt"
)

func main() {
	gormDb, error := db.NewMySqlGormDB()
	if error != nil {
		panic(error)
	}

	only := flag.String("only", "", "Run migrations for a specific model: users, products")
	flag.Parse()

	switch *only {
	case "users":
		error = gormDb.AutoMigrate(&models.User{})
		fmt.Println("Migrated: users table")
	case "products":
		error = gormDb.AutoMigrate(&models.Product{})
		fmt.Println("Migrated: products table")
	case "":
		error = db.RunMigrations(gormDb)
	default:
		fmt.Println("Unknown migration")
	}

	// migrate product table
	// if error := db.RunMigrations(gormDb); error != nil {
	// 	panic(error)
	// }

	if error != nil {
		fmt.Println("Migration failed", error)
		return
	}

	fmt.Println("database migration successfully")
}
