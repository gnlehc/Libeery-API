package main

import (
	"Libeery/database"
	"Libeery/model"
	"Libeery/service"
	"log"
	"os"
)

func main() {
	err := database.DatabaseConnection()

	if err != nil {
		log.Fatalln("Could not connect to database", err)
	}
	database.GlobalDB.AutoMigrate(&model.MsMahasiswa{})
	r := service.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
