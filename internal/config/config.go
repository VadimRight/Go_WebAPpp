package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigDatabase struct {
	Port     string `env:"DB_PORT" env-default:"5432"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Name     string `env:"DB_NAME" env-default:"postgres"`
	User     string `env:"DB_USER" env-default:"user"`
	Password string `env:"DB_DASSWORD"`
}

type HTTPServer struct {
	Port         string        `env:"PORT" env-description:"server port"`
	Timeout      time.Duration `env:"TIMEOUT" env-description:"Timeout"`
	IddleTimeout time.Duration `env:"IDLE_TIMEOUT" env-description:"Idle timeout"`
}

func MustLoad() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exists", configPath)
	}
	var cfg_db ConfigDatabase
	var cfg_server HTTPServer

	if err := cleanenv.ReadConfig(configPath, &cfg_db); err != nil {
		log.Fatal("cannot read database config")
	}
	if err := cleanenv.ReadConfig(configPath, &cfg_server); err != nil {
		log.Fatal("cannot read server config")
	}
}
