package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/service/grpc"
	"gosample/config"
	"gosample/portal/handler"
	"gosample/portal/models"
	"gosample/portal/proto"
)

func init() {
	config.Setup()
	models.Setup()
}
func main() {
	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.portal"),
		micro.Version("latest"),
	)
	// 必须提前初始化
	err := cmd.Init()
	if err != nil {
		log.Fatalf(" cmd init error: %v", err)
	}

	// Initialise service
	service.Init()

	err = proto.RegisterStudentServiceHandler(service.Server(), new(handler.StudentHandler))
	if err != nil {
		log.Fatalf("handler error: %v", err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
