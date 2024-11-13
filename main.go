package main

import (
	"fmt"
	"hotel-api/config"
	"hotel-api/handlers"
	"log"
	"net/http"

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
	fmt.Println("Server berjalan di localhost:8080 dan 192.168.1.4:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler))
}
