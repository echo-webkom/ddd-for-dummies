package model

type Hotel struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Stars        int    `json:"stars"`
	HasBreakfast bool   `json:"hasBreakfast"`
	HasGym       bool   `json:"hasGym"`
	HasPool      bool   `json:"hasPool"`
	Rooms        []Room `json:"rooms"`
}

type Room struct {
	RoomNumber string `json:"roomNumber"`
	Price      int    `json:"price"`
	Capacity   int    `json:"capacity"`
	Occupied   bool   `json:"occupied"`
}
