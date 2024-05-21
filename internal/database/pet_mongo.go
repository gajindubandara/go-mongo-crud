package database

import (
	"context"
	"crud-app/internal/models"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var petCollection *mongo.Collection

func InitPetCollection() {
	petCollection = MongoClient.Database("go-crud").Collection("pets")
}

func CreatePet(pet models.Pet) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return petCollection.InsertOne(ctx, pet)
}

func GetPetByID(id string) (models.Pet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Pet{}, err
	}

	var pet models.Pet
	err = petCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&pet)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// Log or handle "document not found" error
			fmt.Println("Document not found")
		}
		return models.Pet{}, err
	}

	return pet, nil
}

func GetAllPets() ([]models.Pet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := petCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var pets []models.Pet
	if err = cursor.All(ctx, &pets); err != nil {
		return nil, err
	}
	return pets, nil
}

func UpdatePet(id string, pet models.Pet) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.M{"$set": pet}
	return petCollection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
}

func DeletePet(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return petCollection.DeleteOne(ctx, bson.M{"_id": objectID})
}
