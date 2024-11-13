package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"hotel-api/models"

	"github.com/gorilla/mux"
)

var bookings = make(map[int]models.BookingConfirmation)
var bookingID = 1

func ConfirmBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.BookingConfirmation
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Simpan booking ke dalam variabel di memori
	booking.ID = bookingID
	bookings[bookingID] = booking
	bookingID++

	// Kirim respons konfirmasi ke frontend
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Booking              models.BookingConfirmation `json:"booking"`
		ConfirmationMessage  string                    `json:"confirmationMessage"`
	}{
		Booking:             booking,
		ConfirmationMessage: "Booking confirmed! Booking ID: " + strconv.Itoa(booking.ID) + ", Name: " + booking.Name + ", Total Cost: $" + strconv.FormatFloat(booking.TotalCost, 'f', 2, 64),
	}
	json.NewEncoder(w).Encode(response)
}

func GetConfirmedBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		log.Println("Invalid booking ID:", err)
		return
	}

	booking, exists := bookings[id]
	if !exists {
		http.Error(w, "Booking not found", http.StatusNotFound)
		log.Printf("Booking ID %d not found\n", id)
		return
	}

	// Log the booking data field by field
	log.Println("\n--- Confirmed Booking Data Sent ---")
	log.Printf("Booking ID: %d\n", booking.ID)
	log.Printf("Room ID: %d\n", booking.RoomId)
	log.Printf("Room Type: %s\n", booking.RoomType)
	log.Printf("Check-in Date: %s\n", booking.CheckIn)
	log.Printf("Check-out Date: %s\n", booking.CheckOut)
	log.Printf("Adults: %d\n", booking.Adults)
	log.Printf("Children: %d\n", booking.Children)
	log.Printf("Name: %s\n", booking.Name)
	log.Printf("Phone: %s\n", booking.Phone)
	log.Printf("Email: %s\n", booking.Email)
	log.Printf("Total Cost: $%.2f\n", booking.TotalCost)
	log.Printf("Days: %d\n", booking.Days)
	log.Println("-------------------------------")

	// Send the confirmed booking data to frontend
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(booking)
}

// GetAllConfirmedBookings mengambil semua booking yang sudah dikonfirmasi
func GetAllConfirmedBookings(w http.ResponseWriter, r *http.Request) {
	// Ambil semua booking dari map bookings
	var allBookings []models.BookingConfirmation
	for _, booking := range bookings {
		allBookings = append(allBookings, booking)
	}

	// Log each booking field by field
	log.Println("\n--- All Confirmed Bookings Sent ---")
	for _, booking := range allBookings {
		log.Printf("Booking ID: %d\n", booking.ID)
		log.Printf("Room ID: %d\n", booking.RoomId)
		log.Printf("Room Type: %s\n", booking.RoomType)
		log.Printf("Check-in Date: %s\n", booking.CheckIn)
		log.Printf("Check-out Date: %s\n", booking.CheckOut)
		log.Printf("Adults: %d\n", booking.Adults)
		log.Printf("Children: %d\n", booking.Children)
		log.Printf("Name: %s\n", booking.Name)
		log.Printf("Phone: %s\n", booking.Phone)
		log.Printf("Email: %s\n", booking.Email)
		log.Printf("Total Cost: $%.2f\n", booking.TotalCost)
		log.Printf("Days: %d\n", booking.Days)
		log.Println("-------------------------------")
	}
	log.Println("-------------------------------")

	// Send response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allBookings)
}
