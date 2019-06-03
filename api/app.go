package main

import (
	"api/pkg/gdb"
	"api/pkg/gredis"
	"api/pkg/logging"
	"api/pkg/setting"
	"api/routers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	// read .ini files
	setting.Setup()
	// init log system
	logging.Setup(
		setting.AppSetting.GCProjectId,
		setting.AppSetting.GCPStackLogName)
	// init db
	gdb.Setup(
		setting.DBSetting.User,
		setting.DBSetting.Password,
		setting.DBSetting.Host,
		setting.DBSetting.Name)
	// init redis
	err := gredis.Setup(
		setting.RediSetting.Host,
		setting.RediSetting.Password,
		setting.RediSetting.MaxIdle,
		setting.RediSetting.MaxActive,
		setting.RediSetting.IdleTimeout,
	)
	if err != nil {
		logging.Fatal("init api server err in redis: %v ", err)
	}
	jwt.Init(setting.AppSetting.JwtSecret)
}

func main() {
	gin.SetMode(setting.AppSetting.RunMode)

	routersInit := routers.InitRouter(setting.AppSetting.APIPrefix)
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
