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

	// Log the received payload field by field
	log.Println("\n--- Received Request Payload ---")
	log.Printf("Check-in Date: %s\n", payload.CheckIn)
	log.Printf("Check-out Date: %s\n", payload.CheckOut)
	log.Printf("Adults: %d\n", payload.Adults)
	log.Printf("Children: %d\n", payload.Children)
	log.Println("-----------------------------")

	// Get available rooms based on search criteria
	availableRooms := services.FilterAvailableRooms(payload)

	// Check if no rooms are available
	if len(availableRooms) == 0 {
		log.Println("\n--- No Rooms Found ---")
		log.Println("No rooms found for the provided criteria")
		log.Println("-----------------------")
		http.Error(w, "Kamar tidak ditemukan", http.StatusNotFound)
		return
	}

	// Log the response field by field
	log.Println("\n--- Sending Response ---")
	for _, room := range availableRooms {
		log.Printf("Room ID: %d\n", room.ID)
		log.Printf("Room Type: %s\n", room.Type)
		log.Printf("Room Description: %s\n", room.Description)
		log.Printf("Price Per Night: $%.2f\n", room.PricePerNight)
		log.Printf("Total Price: $%.2f\n", room.TotalPrice)
		log.Printf("Capacity: %d\n", room.Capacity)
		log.Printf("Stay Duration: %d nights\n", room.StayDuration)
		log.Printf("Room Image: %s\n", room.Image)
		log.Println("-----------------------")
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(availableRooms)
}
