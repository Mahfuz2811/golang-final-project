package main

import (
	"final-golang-project/db"
	"final-golang-project/models"
	"flag"
	"fmt"
)

func main() {

	only := flag.String("only", "", "Run migration for a specific model: users, products")
	flag.Parse()

	dbConn, err := db.NewMySQLGormDB()
	if err != nil {
		panic(fmt.Sprintf("DB connection failed: %v", err))
	}

	// if err := db.MigrateProductTable(dbConn); err != nil {
	// 	panic(fmt.Sprintf("Migration failed: %v", err))
	// }

	// fmt.Println("Product table migration complete")

	// if err := db.RunMigrations(dbConn); err != nil {
	// 	panic(fmt.Sprintf("Migration failed: %v", err))
	// }

	// fmt.Println("All tables migrated successfully.")

	switch *only {
	case "users":
		err = dbConn.AutoMigrate(&models.User{})
		fmt.Println("Migrated: users")

	case "products":
		err = dbConn.AutoMigrate(&models.Product{})
		fmt.Println("Migrated: products")

	case "":
		err = db.RunMigrations(dbConn)
		fmt.Println("Migrated: all models")

	default:
		fmt.Printf("Unknown migration target: %s\n", *only)
		return
	}

	if err != nil {
		panic(fmt.Sprintf("Migration failed: %v", err))
	}
}
