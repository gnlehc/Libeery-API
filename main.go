package main

import (
	"Libeery/database"
	"Libeery/database/migrate"
	"Libeery/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Establish database connection
	err := database.DatabaseConnection()
	if err != nil {
		log.Fatalln("Could not connect to database:", err)
	}

	// Migrate database tables
	err = migrate.DatabaseMigration(database.GlobalDB)
	if err != nil {
		log.Fatalln("Failed to migrate database tables:", err)
	}
	r := gin.Default()

	service.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
