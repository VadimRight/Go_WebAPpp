package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres_Port     string        `env:"DB_PORT" env-default:"5432"`
	Postgres_Host     string        `env:"DB_HOST" env-default:"localhost"`
	Postgres_Name     string        `env:"DB_NAME" env-default:"postgres"`
	Postgres_User     string        `env:"DB_USER" env-default:"user"`
	Postgres_Password string        `env:"DB_DASSWORD"`
	Server_Port       string        `env:"port" env-description:"server port"`
	Timeout           time.Duration `env:"timeout" env-description:"timeout"`
	IddleTimeout      time.Duration `env:"idle_timeout" env-description:"idle timeout"`
}

type ConfigDatabase struct {
	Port     string `env:"DB_PORT" env-default:"5432"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Name     string `env:"DB_NAME" env-default:"postgres"`
	User     string `env:"DB_USER" env-default:"user"`
	Password string `env:"DB_DASSWORD"`
}

type HTTPServer struct {
	Port         string        `env:"port" env-description:"server port"`
	Timeout      time.Duration `env:"timeout" env-description:"timeout"`
	IddleTimeout time.Duration `env:"idle_timeout" env-description:"idle timeout"`
}

func MustLoad() ConfigDatabase {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exists", configPath)
	}
	var cfg ConfigDatabase

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("cannot read database config")
	}
	return cfg
}
