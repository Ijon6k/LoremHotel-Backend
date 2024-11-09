package handlers

import (
	"encoding/json"
	"hotel-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)
func GetRoomByID(w http.ResponseWriter, r *http.Request) {
	roomIDStr := mux.Vars(r)["id"]
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	// Cari kamar berdasarkan ID
	for _, room := range models.Rooms {
		if room.ID == roomID {
			json.NewEncoder(w).Encode(room)
			return
		}
	}

	http.Error(w, "Room not found", http.StatusNotFound)
}
