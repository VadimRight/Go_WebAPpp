package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/VadimRight/Go_WebApp/internal/config"
	"github.com/VadimRight/Go_WebApp/internal/lib/logger/handlers/slogpretty"
	"github.com/VadimRight/Go_WebApp/internal/server"
	mwlogger "github.com/VadimRight/Go_WebApp/internal/server/middleware/logger"
	"github.com/VadimRight/Go_WebApp/internal/storage/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	db := postgres.InitDB()
	fmt.Println(db)
	test_add := postgres.TestAddUrl()
	fmt.Println(test_add)
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(mwlogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	http.HandleFunc("/", server.GetRoot)
	http.HandleFunc("/hello", server.GetHello)
	port := fmt.Sprintf(":%s", cfg.Server_Port)
	err := http.ListenAndServe(port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = setupPrettySlog()
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

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
