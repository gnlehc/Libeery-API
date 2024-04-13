package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func DatabaseConnection() (err error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("PGHOST")
	dbUsername := os.Getenv("PGUSER")
	dbPassword := os.Getenv("PGPASSWORD")
	dbDatabase := os.Getenv("PGDATABASE")
	dbPort := os.Getenv("PGPORT")

	fmt.Println("Database host:", dbHost)
	fmt.Println("Database username:", dbUsername)
	fmt.Println("Database password:", dbPassword)
	fmt.Println("Database database:", dbDatabase)
	fmt.Println("Database port:", dbPort)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost,
		dbUsername,
		dbPassword,
		dbDatabase,
		dbPort)

	GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	return err
}
