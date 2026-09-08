package hotelapi

import "errors"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

type Hotel struct {
	ID     int
	Name   string
	Closed bool

	info  HotelInfo
	rooms []Room
}

type HotelInfo struct {
	Stars        int
	HasBreakfast bool
	HasGym       bool
	HasPool      bool
}

type Room struct {
	RoomNumber string
	Price      int
	Capacity   int
	Occupied   bool
}

func (c *Client) AllHotels() ([]Hotel, error) {
	h := []Hotel{}
	for _, hh := range hotelsById {
		h = append(h, hh)
	}
	return h, nil
}

func (c *Client) GetHotelById(id int) (Hotel, error) {
	h, ok := hotelsById[id]
	if !ok {
		return Hotel{}, errors.New("not found")
	}
	return h, nil
}

func (c *Client) GetHotelRooms(hotelId int) ([]Room, error) {
	h, ok := hotelsById[hotelId]
	if !ok {
		return nil, errors.New("not found")
	}
	return h.rooms, nil
}

func (c *Client) GetHotelInfo(hotelId int) (HotelInfo, error) {
	h, ok := hotelsById[hotelId]
	if !ok {
		return HotelInfo{}, errors.New("not found")
	}
	return h.info, nil
}

func (c *Client) BookHotelRoom(hotelId int, roomNr string) error {
	h, ok := hotelsById[hotelId]
	if !ok {
		return errors.New("not found")
	}

	for i, r := range h.rooms {
		if r.RoomNumber == roomNr {
			if r.Occupied {
				return errors.New("room already occupied")
			}
			h.rooms[i].Occupied = true
			return nil
		}
	}
	return errors.New("room not found")
}

var hotelsById = map[int]Hotel{
	1: {
		ID:     1,
		Name:   "Grand Plaza Hotel",
		Closed: false,
		info: HotelInfo{
			Stars:        5,
			HasBreakfast: true,
			HasGym:       true,
			HasPool:      true,
		},
		rooms: []Room{
			{RoomNumber: "101", Price: 250, Capacity: 2, Occupied: false},
			{RoomNumber: "102", Price: 300, Capacity: 3, Occupied: true},
			{RoomNumber: "201", Price: 450, Capacity: 4, Occupied: false},
			{RoomNumber: "301", Price: 600, Capacity: 2, Occupied: true},
			{RoomNumber: "302", Price: 750, Capacity: 5, Occupied: false},
		},
	},

	2: {
		ID:     2,
		Name:   "Seaside Resort",
		Closed: false,
		info: HotelInfo{
			Stars:        4,
			HasBreakfast: true,
			HasGym:       true,
			HasPool:      true,
		},
		rooms: []Room{
			{RoomNumber: "10", Price: 180, Capacity: 2, Occupied: false},
			{RoomNumber: "11", Price: 220, Capacity: 4, Occupied: true},
			{RoomNumber: "20", Price: 350, Capacity: 4, Occupied: false},
		},
	},

	3: {
		ID:     3,
		Name:   "City Center Inn",
		Closed: false,
		info: HotelInfo{
			Stars:        3,
			HasBreakfast: true,
			HasGym:       false,
			HasPool:      false,
		},
		rooms: []Room{
			{RoomNumber: "1", Price: 90, Capacity: 1, Occupied: false},
			{RoomNumber: "2", Price: 120, Capacity: 2, Occupied: true},
		},
	},

	4: {
		ID:     4,
		Name:   "Mountain View Lodge",
		Closed: false,
		info: HotelInfo{
			Stars:        4,
			HasBreakfast: true,
			HasGym:       false,
			HasPool:      true,
		},
		rooms: []Room{
			{RoomNumber: "A1", Price: 140, Capacity: 2, Occupied: false},
			{RoomNumber: "A2", Price: 160, Capacity: 3, Occupied: false},
			{RoomNumber: "B1", Price: 280, Capacity: 4, Occupied: true},
			{RoomNumber: "B2", Price: 320, Capacity: 5, Occupied: false},
		},
	},

	5: {
		ID:     5,
		Name:   "Budget Stay",
		Closed: false,
		info: HotelInfo{
			Stars:        2,
			HasBreakfast: false,
			HasGym:       false,
			HasPool:      false,
		},
		rooms: []Room{
			{RoomNumber: "5", Price: 50, Capacity: 1, Occupied: false},
		},
	},

	6: {
		ID:     6,
		Name:   "Royal Palace",
		Closed: false,
		info: HotelInfo{
			Stars:        5,
			HasBreakfast: true,
			HasGym:       true,
			HasPool:      true,
		},
		rooms: []Room{
			{RoomNumber: "101", Price: 400, Capacity: 2, Occupied: true},
			{RoomNumber: "102", Price: 450, Capacity: 3, Occupied: false},
			{RoomNumber: "201", Price: 700, Capacity: 4, Occupied: true},
			{RoomNumber: "202", Price: 900, Capacity: 5, Occupied: false},
			{RoomNumber: "301", Price: 1200, Capacity: 2, Occupied: false},
		},
	},

	7: {
		ID:     7,
		Name:   "Forest Retreat",
		Closed: false,
		info: HotelInfo{
			Stars:        3,
			HasBreakfast: true,
			HasGym:       false,
			HasPool:      true,
		},
		rooms: []Room{
			{RoomNumber: "C1", Price: 130, Capacity: 2, Occupied: false},
			{RoomNumber: "C2", Price: 150, Capacity: 2, Occupied: true},
			{RoomNumber: "D1", Price: 200, Capacity: 4, Occupied: false},
		},
	},

	8: {
		ID:     8,
		Name:   "Airport Hotel",
		Closed: false,
		info: HotelInfo{
			Stars:        3,
			HasBreakfast: true,
			HasGym:       true,
			HasPool:      false,
		},
		rooms: []Room{
			{RoomNumber: "201", Price: 110, Capacity: 1, Occupied: true},
			{RoomNumber: "202", Price: 130, Capacity: 2, Occupied: false},
			{RoomNumber: "203", Price: 160, Capacity: 3, Occupied: false},
			{RoomNumber: "204", Price: 190, Capacity: 4, Occupied: true},
		},
	},

	9: {
		ID:     9,
		Name:   "Old Town Hotel",
		Closed: true,
		info: HotelInfo{
			Stars:        2,
			HasBreakfast: false,
			HasGym:       false,
			HasPool:      false,
		},
		rooms: []Room{
			{RoomNumber: "7", Price: 70, Capacity: 1, Occupied: false},
			{RoomNumber: "8", Price: 90, Capacity: 2, Occupied: false},
		},
	},

	10: {
		ID:     10,
		Name:   "Lakeside Hotel",
		Closed: false,
		info: HotelInfo{
			Stars:        4,
			HasBreakfast: true,
			HasGym:       true,
			HasPool:      false,
		},
		rooms: []Room{
			{RoomNumber: "101", Price: 170, Capacity: 2, Occupied: false},
			{RoomNumber: "102", Price: 190, Capacity: 2, Occupied: true},
			{RoomNumber: "201", Price: 250, Capacity: 3, Occupied: false},
			{RoomNumber: "202", Price: 300, Capacity: 4, Occupied: false},
			{RoomNumber: "301", Price: 500, Capacity: 5, Occupied: true},
		},
	},
}
