package models

type BookingConfirmation struct {
	ID        int     `json:"id"`
	RoomId    int     `json:"roomId"`
	RoomType  string  `json:"roomType"`
	CheckIn   string  `json:"checkIn"`
	CheckOut  string  `json:"checkOut"`
	Adults    int     `json:"adults"`
	Children  int     `json:"children"`
	Name      string  `json:"name"`
	Phone     string  `json:"phone"`
	Email     string  `json:"email"`
	TotalCost float64 `json:"totalCost"`
	Days      int     `json:"days"`
}
