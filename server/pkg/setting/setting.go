package setting

import (
	"log"

	"github.com/go-ini/ini"
)

var AppSetting = &App{}
var ServerSetting = &Server{}
var RediSetting = &Redis{}
var DBSetting = &Database{}

func Setup() {
	configMap("configs/server.ini", ServerSetting)
	configMap("configs/app.ini", AppSetting)
	configMap("configs/db.ini", DBSetting)
	configMap("configs/redis.ini", RediSetting)
}

func configMap(filename string, v interface{}) {
	cfg, err := ini.Load(filename)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse '%s': %v", filename, err)
	}

	err = cfg.MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo Setting err: %v", err)
	}
}
