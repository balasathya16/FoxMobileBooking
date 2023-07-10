package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/balasathya16/FoxBooking/db"
	"github.com/balasathya16/FoxBooking/models"

	"github.com/gorilla/mux"
)

func CreateCricketCourt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var court models.CricketCourt
	_ = json.NewDecoder(r.Body).Decode(&court)

	// Save the court to the database
	database, err := db.ConnectDB()
	if err != nil {
		// Handle the error appropriately
	}

	// Get the collection
	collection := database.Collection("cricket_courts")

	// Insert the court document
	_, err = collection.InsertOne(context.TODO(), court)
	if err != nil {
		// Handle the error appropriately
	}

	json.NewEncoder(w).Encode(court)
}

func GetCricketCourt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	courtID := params["id"]

	// Retrieve the cricket court from the database using MongoDB driver based on courtID
	// Implement the necessary logic to fetch the court details and assign it to the `court` variable

	court := models.CricketCourt{
		ID:       courtID,
		Location: "Sample Location",
		// Set other properties based on the fetched data
	}

	json.NewEncoder(w).Encode(court)
}
