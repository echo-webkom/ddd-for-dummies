package port

import "ddd-for-dummies/domain/model"

type HotelRepo interface {
	GetAllHotels() ([]model.Hotel, error)
}
