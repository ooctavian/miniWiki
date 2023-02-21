package miniWiki

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Logger   LoggerConfig
}

type DatabaseConfig struct {
	DatabaseUrl string `env:"DATABASE_URL,required"`
	ImageDir    string `env:"IMAGE_DIR" envDefault:"images/"`
}

type ServerConfig struct {
	Port string `env:"PORT" envDefault:"4444"`
}

type LoggerConfig struct {
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
