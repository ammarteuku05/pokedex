package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	DbUser                  string `envconfig:"DB_USER" required:"true" default:"user"`
	DbPass                  string `envconfig:"DB_PASS" required:"true" default:"password"`
	DbName                  string `envconfig:"DB_NAME" required:"true" default:"db"`
	DbHost                  string `envconfig:"DB_HOST" required:"true" default:"127.0.0.1"`
	DbPort                  int    `envconfig:"DB_PORT" required:"true" default:"3306"`
	JwtSecret               string `envconfig:"JWT_SECRET_KEY" required:"true" default:"secret"`
	DbMaxOpenConnection     int    `envconfig:"DB_MAX_OPEN_CONNECTION" required:"true" default:"32"`
	DbMaxIdleConnection     int    `envconfig:"DB_MAX_IDLE_CONNECTION" required:"true" default:"8"`
	DbMaxLifeTimeConnection string `envconfig:"DB_MAX_LIFE_TIME_CONNECTION" required:"true" default:"1h"`
}

var cfg Configuration

func init() {
	_ = envconfig.Process("", &cfg)
}

func GetConfig() Configuration {
	return cfg
}
