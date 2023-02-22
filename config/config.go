package config

import (
	"time"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Logger   LoggerConfig
}

type DatabaseConfig struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
	ImageDir    string `env:"IMAGE_DIR" envDefault:"images/"`
}

type ServerConfig struct {
	Port    string        `env:"PORT" envDefault:"4444"`
	Timeout time.Duration `env:"Timeout" envDefault:"10s"`
}

type LoggerConfig struct {
	Level string `env:"LOG_LEVEL" envDefault:"Info"`
}

func InitConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	var config Config
	err = env.Parse(&config)
	if err != nil {
		return nil, err
	}
	return &config, err
}
