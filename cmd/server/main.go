package main

import (
	"crud-app/config"
	"crud-app/internal/database"
	"crud-app/internal/routes"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()
	database.InitMongoDB()
	database.InitPetCollection()

	router := routes.SetupRouter()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
