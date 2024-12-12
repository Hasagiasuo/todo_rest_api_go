package configs

import (
	"log/slog"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string 		`yaml:"env"`
	ServerConfig	`yaml:"server"`
	DBConfig			`yaml:"db"`
}

type DBConfig struct {
	Host 			string 		`yaml:"host"`
  Port 			string 		`yaml:"port"`
  Name 			string 		`yaml:"name"`
  DB_name 	string 		`yaml:"db_name"`
  SSL_mode 	string 		`yaml:"ssl_mode"`
}

type ServerConfig struct {
	Address 				string 		`yaml:"address"`
	Timeout 				float32 	`yaml:"timeout"`
	Iddle_timeout		float32		`yaml:"iddle_timeout"`
}

func InitServerConfig() *Config {
	var cfg Config
	if err := cleanenv.ReadConfig("./config/conf.yaml", &cfg); err != nil {
		slog.Info("Config file not found")
	}
	return &cfg
}