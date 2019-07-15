package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-web"
	"gosample/api/routers"
	"gosample/config"
	"gosample/portal/models"
	"log"
	"net/http"
)

var err error

func init() {
	config.Setup()
	models.Setup()
}
func main() {

	// 创建 micro 服务
	service := web.NewService(
		web.Name("go.micro.api.sample"),
	)
	gin.SetMode(config.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := config.ServerSetting.ReadTimeout
	writeTimeout := config.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	// 用 gin 注册go-micro handler
	service.Handle("/", routersInit)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("run gin error: %s", err)
	}
	if err := service.Run(); err != nil {
		log.Fatalf("run go-micro error: %s", err)
	}
}
