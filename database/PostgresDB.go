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
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}
	// env live railway
	dbUsername := os.Getenv("PGUSER")
	dbPassword := os.Getenv("PGPASSWORD")
	dbHost := os.Getenv("PGHOST")
	dbPort := os.Getenv("PGPORT")
	dbDatabase := os.Getenv("PGDATABASE")

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
