package main

import (
	geo "traffic-dispatcher/proto/geo"
	"traffic-dispatcher/service/lbs/handler"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
)

func main() {
	// New Service
	service := micro.NewService(
		// micro.Registry(reg),
		micro.Name("go.micro.srv.lbs"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	geo.RegisterGeoLocationHandler(service.Server(), new(handler.GeoLocation))

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
