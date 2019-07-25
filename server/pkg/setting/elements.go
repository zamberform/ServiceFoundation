package setting

type App struct {
	RunMode string

	JwtSecret string
	PrefixUrl string
	APIPrefix string
	CMSPrefix string

	RuntimeRootPath string
	TimeFormat      string

	GCProjectId     string
	GCPStackLogName string
}

type Server struct {
	HttpPort     int
	ReadTimeout  int
	WriteTimeout int
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
