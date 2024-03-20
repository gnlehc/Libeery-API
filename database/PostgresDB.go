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
	// env live railway
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	fmt.Println(dbUsername)
	fmt.Println(dbDatabase)
	fmt.Println(dbHost)
	fmt.Println(dbPort)
	fmt.Println(dbPassword)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbDatabase)

	GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	return err
}
