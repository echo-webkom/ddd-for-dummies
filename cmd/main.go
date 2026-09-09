package main

import (
	"ddd-for-dummies/adapter/http"
	"ddd-for-dummies/domain/service"
	"ddd-for-dummies/infra/hotelapi"
)

func main() {
	hotelRepo := hotelapi.NewHotelRepo()
	hotelSercice := service.NewHotelService(hotelRepo)

	http.RunServer(hotelSercice)
}
