package main

import (
	"os"

	"log/slog"

	"github.com/VadimRight/Go_WebApp/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(envLocal)
	log.Info(
		"starting url-shortener",
		slog.String("env", envLocal),
		slog.String("version", "123"),
	)
	log.Info(
		"this is log",
		slog.String("Postgres Name: ", cfg.Postgres_Name),
		slog.String("Postgres Port: ", cfg.Postgres_Port),
		slog.String("Postgres Host: ", cfg.Postgres_Host),
		slog.String("Postgres User: ", cfg.Postgres_User),
		slog.String("Server Port: ", cfg.Server_Port),
		slog.Duration("Timeout: ", cfg.Timeout),
		slog.Duration("Idle Timeout: ", cfg.IdleTimeout),
	)
}
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
