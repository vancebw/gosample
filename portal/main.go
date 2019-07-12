package portal

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"gosample/config"
	"gosample/portal/models"
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

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
