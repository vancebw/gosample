package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"gosample/config"
	"gosample/portal/handler"
	"gosample/portal/models"
	portal "gosample/portal/proto"
)

func init() {
	config.Setup()
	models.Setup()
}
func main() {
	// New Service
	service := micro.NewService(
		micro.Name("gosample.srv.portal"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	err := portal.RegisterStudentServiceHandler(service.Server(), new(handler.StudentHandler))
	if err != nil {
		log.Fatalf("handler error: %v", err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
