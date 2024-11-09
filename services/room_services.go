package services

import (
	"hotel-api/models"
	"time"
)

// FilterAvailableRooms memfilter kamar berdasarkan kapasitas dan tanggal check-in dan check-out
func FilterAvailableRooms(payload models.SearchPayload) []models.Room {
	var available []models.Room

	// Parse tanggal check-in dan check-out
	checkInDate, _ := time.Parse("2006-01-02", payload.CheckIn)
	checkOutDate, _ := time.Parse("2006-01-02", payload.CheckOut)
	days := int(checkOutDate.Sub(checkInDate).Hours() / 24)

	// Filter kamar berdasarkan kapasitas dan hitung harga total
	for _, room := range models.Rooms {
		if payload.Adults+payload.Children <= room.Capacity {
			available = append(available, models.Room{
				Type:          room.Type,
				ID:            room.ID,
				Description:   room.Description,
				PricePerNight: room.PricePerNight,
				TotalPrice:    room.PricePerNight * float64(days),
				Capacity:      room.Capacity,
				StayDuration:  days,
				Image: 		   room.Image,
			})
		}
	}

	return available
}
