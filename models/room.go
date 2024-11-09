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
		Image:         "https://cdn.discordapp.com/attachments/1138842520172904611/1304290963543556167/image_8.png?ex=672edb1e&is=672d899e&hm=3881713cc7e1d25e4a0291839366465b0c5399c0ee1f1eaafcdbf0532b3e4a77&",
	},
	{
		ID:            2,
		Type:          "Deluxe",
		Description:   "Spacious room with a beautiful view and additional amenities.",
		PricePerNight: 300,
		TotalPrice:    0,
		Capacity:      2,
		StayDuration:  0,
		Image:         "https://cdn.discordapp.com/attachments/1138842520172904611/1304290963862192148/image_14.png?ex=672edb1e&is=672d899e&hm=77fe966087b137991728cba7466637dd6cca6e61022a66787ea36deec8e8c1be&",
	},
	{
		ID:            3,
		Type:          "Suite",
		Description:   "Luxurious suite with a living area, balcony, and premium facilities.",
		PricePerNight: 900,
		TotalPrice:    0,
		Capacity:      4,
		StayDuration:  0,
		Image:         "https://media.discordapp.net/attachments/1138842520172904611/1304290963203686441/image_9.png?ex=672edb1e&is=672d899e&hm=23b1cf00fd769f19b138c13a3ae3d04740ed3f0c5e13f8998c23974dfdd9aca5&=&format=webp&quality=lossless&width=778&height=473",
	},
	{
		ID:            4,
		Type:          "Family Room",
		Description:   "Spacious room perfect for families, includes multiple beds and extra space.",
		PricePerNight: 700,
		TotalPrice:    0,
		Capacity:      5,
		StayDuration:  0,
		Image:         "https://media.discordapp.net/attachments/1138842520172904611/1304290964181094420/image_11.png?ex=672edb1e&is=672d899e&hm=f016f59a2e3b2232be6d5848d88f99a1d495bb9994cf758236cd28f5ba12c119&=&format=webp&quality=lossless&width=762&height=473",
	},
}
