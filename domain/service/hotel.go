package service

import (
	"ddd-for-dummies/domain/model"
	"ddd-for-dummies/domain/port"
)

type HotelService struct {
	hotelRepo port.HotelRepo
}

func NewHotelService(hotelRepo port.HotelRepo) *HotelService {
	return &HotelService{hotelRepo}
}

func (h *HotelService) GetAllHotels() ([]model.Hotel, error) {
	return h.hotelRepo.GetAllHotels()
}
