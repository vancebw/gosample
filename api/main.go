package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"gosample/api/routers"
	"gosample/config"
	"log"
	server "net/http"
)

var err error

func init() {
	config.Setup()
}
func main() {
	// 创建 micro 服务
	service := web.NewService(
		web.Name("gosample.api"),
	)
	service.Init()
	gin.SetMode(config.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := config.ServerSetting.ReadTimeout
	writeTimeout := config.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &server.Server{
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
