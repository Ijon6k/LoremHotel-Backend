package main

import (
	"fmt"
	"log"
	"net/http"

	"hotel-api/config"
	"hotel-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Setup CORS middleware
	handler := config.SetupCORS().Handler(r)

	// Routes
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/search", handlers.SearchHandler).Methods("POST")
	r.HandleFunc("/booking/{id:[0-9]+}", handlers.GetRoomByID).Methods("GET")
	r.HandleFunc("/confirm", handlers.ConfirmBooking).Methods("POST")
	r.HandleFunc("/confirmed-booking/{id:[0-9]+}", handlers.GetConfirmedBooking).Methods("GET")
	r.HandleFunc("/confirmed-bookings", handlers.GetAllConfirmedBookings).Methods("GET")



	// Start server
	fmt.Println("Server berjalan di port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
