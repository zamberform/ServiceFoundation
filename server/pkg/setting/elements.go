package setting

import "time"

type App struct {
	RunMode string

	JwtSecret string
	PrefixUrl string
	APIPrefix string

	RuntimeRootPath string
	TimeFormat      string

	GCProjectId     string
	GCPStackLogName string
}

type Server struct {
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}
