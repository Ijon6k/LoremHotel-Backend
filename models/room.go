package models

type SearchPayload struct {
	CheckIn  string `json:"checkin"`
	CheckOut string `json:"checkout"`
	Adults   int    `json:"adults"`
	Children int    `json:"children"`
}

type Room struct {
	ID            int     `json:"id"`
	Type          string  `json:"type"`
	Description   string  `json:"description"`
	PricePerNight float64 `json:"price_per_night"`
	TotalPrice    float64 `json:"total_price"`
	Capacity      int     `json:"capacity"`
	StayDuration  int     `json:"stay_duration"`
	Image         string  `json:"image"` // Field image untuk URL gambar
}

// Sample data for available rooms
var Rooms = []Room{
	{
		ID:            1,
		Type:          "Standard",
		Description:   "A comfortable room with basic amenities.",
		PricePerNight: 100,
		TotalPrice:    0,
		Capacity:      2,
		StayDuration:  0,
		Image:         "/image/standard.png",
	},
	{
		ID:            2,
		Type:          "Deluxe",
		Description:   "Spacious room with a beautiful view and additional amenities.",
		PricePerNight: 300,
		TotalPrice:    0,
		Capacity:      2,
		StayDuration:  0,
		Image:         "/image/deluxe.png",
	},
	{
		ID:            3,
		Type:          "Suite",
		Description:   "Luxurious suite with a living area, balcony, and premium facilities.",
		PricePerNight: 900,
		TotalPrice:    0,
		Capacity:      4,
		StayDuration:  0,
		Image:         "/image/suite.png",
	},
	{
		ID:            4,
		Type:          "Family Room",
		Description:   "Spacious room perfect for families, includes multiple beds and extra space.",
		PricePerNight: 700,
		TotalPrice:    0,
		Capacity:      5,
		StayDuration:  0,
		Image:         "/image/family.png",
	},
}
