package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gosample/config"
	"gosample/models"
	"gosample/routers"
	"log"
	"net/http"
)

var err error

func init() {
	config.Setup()
	models.Setup()
}
func main() {
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

	server.ListenAndServe()
}
