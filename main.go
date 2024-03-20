package main

import (
	"Libeery/database"
	"Libeery/model"
	"Libeery/service"
	"log"
)

func main() {
	err := database.DatabaseConnection()

	if err != nil {
		log.Fatalln("Could not connect to database", err)
	}
	database.GlobalDB.AutoMigrate(&model.MsMahasiswa{})
	r := service.Router()
	r.Run(":8080")
}
