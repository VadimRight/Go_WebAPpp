package config

import (
	"log"
	"os"
	"time"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Postgres_Port     string        `env:"POSTGRES_PORT" env-default:"5432"`
	Postgres_Host     string        `env:"POSTGRES_HOST" env-default:"localhost"`
	Postgres_Name     string        `env:"POSTGRES_NAME" env-default:"godb"`
	Postgres_User     string        `env:"POSTGRES_USER" env-default:"postgres"`
	Postgres_Password string        `env:"POSTGRES_DASSWORD" env-default:"postgres"`
	Server_Port       string        `env:"SERVER_PORT" env-default:"8000"`
	Timeout           time.Duration `env:"TIMEOUT" env-default:"30s"`
	IdleTimeout       time.Duration `env:"IDLE_TIMEOUT" env-default:"110s"`
	Server_Addr string `env:"SERVER_ADDR" env-description:"server adderess"`
	Env string `env:"ENV"`
}

type ConfigDatabase struct {
	Postgres_Port     string `env:"POSTGRES_PORT" env-default:"5432"`
	Postgres_Host     string `env:"POSTGRES_HOST" env-default:"localhost"`
	Postgres_Name     string `env:"POSTGRES_NAME" env-default:"godb"`
	Postgres_User     string `env:"POSTGRES_USER" env-default:"postgres"`
	Postgres_Password string `env:"POSTGRES_DASSWORD"`
}

type HTTPServer struct {
	Server_Addr string `env:"SERVER_ADDR" env-description:"server adderess" env-default:"localhost:8000"`
	Server_Port  string        `env:"SERVER_PORT" env-description:"server port"`
	Timeout      time.Duration `env:"TIMEOUT" env-description:"timeout"`
	IddleTimeout time.Duration `env:"IDLE_TIMEOUT" env-description:"idle timeout"`
}

func MustLoad() Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}		
	envDev := os.Getenv("ENV")
	log := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	log.Printf("ENV is %s", envDev)
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exists", configPath)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("cannot read database config")
	}
	return cfg
}
