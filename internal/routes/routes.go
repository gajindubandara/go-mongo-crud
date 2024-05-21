package routes

import (
	"crud-app/internal/handlers"
	"crud-app/internal/middleware"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.ErrorHandler)
	r.HandleFunc("/healthcheck", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/pets", handlers.CreatePet).Methods("POST")
	r.HandleFunc("/pets/{id}", handlers.GetPetByID).Methods("GET")
	r.HandleFunc("/pets", handlers.GetAllPets).Methods("GET")
	r.HandleFunc("/pets/{id}", handlers.UpdatePet).Methods("PUT")
	r.HandleFunc("/pets/{id}", handlers.DeletePet).Methods("DELETE")
	return r
}
