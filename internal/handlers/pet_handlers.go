package handlers

import (
	"crud-app/internal/database"
	"crud-app/internal/errors"
	"crud-app/internal/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CreatePet(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		panic(errors.BadRequestError("Invalid request payload"))
	}

	// Validate the required fields
	if pet.Name == "" {
		panic(errors.BadRequestError("Name is required"))
	}

	// Validate the date format
	if pet.DateOfBirth != "" {
		if _, err := time.Parse("2006-01-02", pet.DateOfBirth); err != nil {
			panic(errors.BadRequestError("Invalid date format. Use YYYY-MM-DD."))
		}
	}

	// Create the pet in the database
	_, err := database.CreatePet(pet)
	if err != nil {
		panic(errors.InternalServerError(err))
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	response := "Pet added successfully"
	w.Write([]byte(response))
}

func GetPetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pet, err := database.GetPetByID(params["id"])
	if err != nil {
		panic(errors.NotFoundError("Pet not found"))
	}
	json.NewEncoder(w).Encode(pet)
}

func GetAllPets(w http.ResponseWriter, r *http.Request) {
	pets, err := database.GetAllPets()
	if err != nil {
		panic(errors.InternalServerError(err))
	}
	json.NewEncoder(w).Encode(pets)
}

func UpdatePet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var pet models.Pet

	_, err := database.GetPetByID(params["id"])
	if err != nil {
		panic(errors.NotFoundError("Pet not found"))
	}

	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		panic(errors.BadRequestError("Invalid request payload"))
	}

	// Validate the required fields
	if pet.Name == "" {
		panic(errors.BadRequestError("Name is required"))
	}

	// Validate the date format
	if pet.DateOfBirth != "" {
		if _, err := time.Parse("2006-01-02", pet.DateOfBirth); err != nil {
			panic(errors.BadRequestError("Invalid date format. Use YYYY-MM-DD."))
		}
	}

	_, err = database.UpdatePet(params["id"], pet)
	if err != nil {
		panic(errors.InternalServerError(err))
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	response := "Pet updated successfully"
	w.Write([]byte(response))
}

func DeletePet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	_, err := database.GetPetByID(params["id"])
	if err != nil {
		panic(errors.NotFoundError("Pet not found"))
	}

	_, err = database.DeletePet(params["id"])
	if err != nil {
		panic(errors.InternalServerError(err))
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	response := "Pet deleted successfully"
	w.Write([]byte(response))
}
