package routes

import (
	"ddd-for-dummies/domain/service"
	"errors"
	"net/http"

	"github.com/jesperkha/fuse"
)

func GetAllHotels(hotelService *service.HotelService) func(ctx *fuse.Context) error {
	return func(ctx *fuse.Context) error {
		hotels, err := hotelService.GetAllHotels()
		if err != nil {
			return ctx.Error(errors.New("could not get hotels"), http.StatusInternalServerError)
		}

		return ctx.JSON(hotels)
	}
}
