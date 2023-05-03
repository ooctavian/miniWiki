package config

import (
	"time"

	"github.com/caarlos0/env/v7"
)

type Config struct {
	Env             string        `env:"ENV" envDefault:"dev"`
	SessionDuration time.Duration `env:"SESSION_DURATION" envDefault:"1h"`
	Database        DatabaseConfig
	Server          ServerConfig
	Logger          LoggerConfig
	Argon2id        Argon2idConfig
	S3Bucket        S3Config
}

type DatabaseConfig struct {
	DatabaseURL      string `env:"DATABASE_URL,required"`
	ImageDir         string `env:"IMAGE_DIR" envDefault:"images/"`
	ResourceImageDir string `env:"IMAGE_RESOURCE_DIR" envDefault:"resource"`
	ProfileImageDir  string `env:"IMAGE_PROFILE_DIR" envDefault:"profile"`
}

type ServerConfig struct {
	Port    string        `env:"PORT" envDefault:"4444"`
	Timeout time.Duration `env:"TIMEOUT" envDefault:"10s"`
}

type S3Config struct {
	Endpoint  string `env:"AWS_ENDPOINT" envDefault:"s3.eu-central-1.amazonaws.com"`
	Region    string `env:"AWS_REGION" envDefault:"eu-central-1"`
	KeyID     string `env:"AWS_KEY_ID,required"`
	SecretKey string `env:"AWS_SECRET_KEY,required"`
	Token     string `env:"AWS_TOKEN" envDefault:""`
}

type LoggerConfig struct {
	Level     string `env:"LOG_LEVEL" envDefault:"info"`
	Formatter string `env:"LOG_FORMATTER" envDefault:"json"`
}

type Argon2idConfig struct {
	Memory      uint32 `env:"ARGON2ID_MEMORY" envDefault:"65536"`
	Iterations  uint32 `env:"ARGON2ID_ITERATIONS" envDefault:"3"`
	Parallelism uint8  `env:"ARGON2ID_PARALLELISM" envDefault:"2"`
	SaltLength  uint32 `env:"ARGON2ID_SALT_LENGTH" envDefault:"16"`
	KeyLength   uint32 `env:"ARGON2ID_KEY_LENGTH" envDefault:"32"`
}

func InitConfig() (*Config, error) {
	var config Config
	err := env.Parse(&config)
	if err != nil {
		return nil, err
	}
	return &config, err
}
