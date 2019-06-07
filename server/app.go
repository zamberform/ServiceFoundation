package main

import (
	"fmt"
	"net/http"
	"server/middleware/jwt"
	"server/pkg/gdb"
	"server/pkg/gredis"
	"server/pkg/logging"
	"server/pkg/setting"
	"server/routers"
	"time"

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
		setting.DBSetting.Type,
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

	jwt.Setup(setting.AppSetting.JwtSecret)
}

func main() {
	gin.SetMode(setting.AppSetting.RunMode)

	routersInit := routers.InitRouter(setting.AppSetting.APIPrefix)
	readTimeout := time.Duration(setting.ServerSetting.ReadTimeout) * time.Second
	writeTimeout := time.Duration(setting.ServerSetting.WriteTimeout) * time.Second
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
