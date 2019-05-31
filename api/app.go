package main

import (
	"api/pkg/gdb"
	"api/pkg/gredis"
	"api/pkg/setting"
	"api/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	gdb.Setup()
	err := gredis.Setup()
	if err != nil {
		log.Fatalf("init api server err in redis: %v ", err)
	}
}

func main() {
	gin.SetMode(setting.AppSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	apiServer := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	apiServer.ListenAndServe()
}
