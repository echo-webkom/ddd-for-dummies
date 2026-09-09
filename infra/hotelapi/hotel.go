package hotelapi

import (
	"ddd-for-dummies/domain/model"
	"ddd-for-dummies/domain/port"
	"ddd-for-dummies/pkg/hotelapi"
)

type hotelRepo struct {
	client *hotelapi.Client
}

func NewHotelRepo() port.HotelRepo {
	return &hotelRepo{
		client: hotelapi.NewClient(),
	}
}

func (h *hotelRepo) GetAllHotels() (hotels []model.Hotel, err error) {
	apiHotels, err := h.client.AllHotels()
	if err != nil {
		return hotels, err
	}

	for _, hotel := range apiHotels {
		if hotel.Closed {
			continue
		}

		apiRooms, err := h.client.GetHotelRooms(hotel.ID)
		if err != nil || len(apiRooms) == 0 {
			continue
		}

		var rooms []model.Room
		for _, r := range apiRooms {
			rooms = append(rooms, model.Room{
				RoomNumber: r.RoomNumber,
				Price:      r.Price,
				Capacity:   r.Capacity,
				Occupied:   r.Occupied,
			})
		}

		info, err := h.client.GetHotelInfo(hotel.ID)
		if err != nil {
			return hotels, err
		}

		hotels = append(hotels, model.Hotel{
			Name:         hotel.Name,
			Rooms:        rooms,
			Stars:        info.Stars,
			HasBreakfast: info.HasBreakfast,
			HasGym:       info.HasGym,
			HasPool:      info.HasPool,
		})
	}

	return hotels, err
}
