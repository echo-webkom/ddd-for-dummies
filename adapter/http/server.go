package http

import (
	"ddd-for-dummies/adapter/http/routes"
	"ddd-for-dummies/domain/service"
	"net/http"

	"github.com/jesperkha/fuse"
)

func RunServer(hotelService *service.HotelService) {
	mux := fuse.NewMux(
		fuse.LoggerMiddleware(fuse.BasicLogger{}),
	)

	mux.Handle("GET", "/hotels", routes.GetAllHotels(hotelService))

	mux.Handle("GET", "/", func(ctx *fuse.Context) error {
		return ctx.ServeFile("web/index.html")
	})

	http.ListenAndServe(":8080", mux)
}
