package configs

type DbConfig struct {
	HandlerName string
	Host        string
	Port        string
	Dbuser      string
	Dbpassword  string
	Database    string
}

func LoadDbConfig()
