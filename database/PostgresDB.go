package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func DatabaseConnection() (err error) {
	// env ada di terminal linux
	// dbUsername := os.Getenv("DB_USERNAME")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbDatabase := os.Getenv("DB_DATABASE")
	// ambil credential dari docker-compose
	dbHost := "localhost"
	dbUsername := "admin"
	dbDatabase := "Libeery-DB"
	dbPassword := "secret"
	dbPort := "5444"

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
