package handlers

import (
	"encoding/json"
	"hotel-api/models"
	"hotel-api/services"
	"log"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.SearchPayload

	// Parse request body
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("Failed to parse request payload:", err)
		return
	}

	// Log the received payload
	log.Printf("Received request payload: %+v\n", payload)

	// Get available rooms based on search criteria
	availableRooms := services.FilterAvailableRooms(payload)

	// Check if no rooms are available
	if len(availableRooms) == 0 {
		log.Println("No rooms found for the provided criteria")
		http.Error(w, "Kamar tidak ditemukan", http.StatusNotFound)
		return
	}

	// Log the response that will be sent
	log.Printf("Sending response: %+v\n", availableRooms)

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(availableRooms)
}
