package apiserver

type Config struct {
	BindAddr    string
	LogLevel    string
	DatabaseURL string
	SessionKey  string
	ClientUrl	string
}

func NewConfig() *Config {
	return &Config{
		BindAddr:		":8070",
		LogLevel:		"debug",
		SessionKey:		"jdfhdfdj",
		DatabaseURL:	"host=localhost dbname=postgres sslmode=disable port=5432 password=1234 user=d",
		ClientUrl:		"http://127.0.0.1:9000",
		//ClientUrl:		"https://comandus.now.sh",
	}
}
