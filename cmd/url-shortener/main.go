package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/VadimRight/Url-Saver/docs"
	"github.com/VadimRight/Url-Saver/internal/config"
	save "github.com/VadimRight/Url-Saver/internal/server/handler/url"
	"github.com/VadimRight/Url-Saver/internal/lib/logger/handlers/slogpretty"
	mwlogger "github.com/VadimRight/Url-Saver/internal/server/middleware/logger"
	"github.com/VadimRight/Url-Saver/internal/storage/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "github.com/VadimRight/Url-Saver/docs"
)

// @title URLSaver
// @version 1.0
// @license.name Apache 2.0
// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @externalDocs.description OpenAPI
// @externalDocs.url https://swagger.io/resources/open-api/

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	log.Info(
		"starting url-shortener",
		slog.String("env", envDev),
		slog.String("version", "123"),
	)
	log.Info(
		"this is log",
		slog.String("Postgres Name", cfg.Postgres_Name),
		slog.String("Postgres Port", cfg.Postgres_Port),
		slog.String("Postgres Host", cfg.Postgres_Host),
		slog.String("Postgres User", cfg.Postgres_User),
		slog.String("Server Port", cfg.Server_Port),
		slog.Duration("Timeout", cfg.Timeout),
		slog.Duration("Idle Timeout", cfg.IdleTimeout),
	)
	daba := postgres.GORMStorage{}
	db, err := daba.InitDB()
	if err != nil {
		log.Error("failed to init storage", err)
		os.Exit(1)
	}
	fmt.Println(db)
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(mwlogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Post("/new_url", save.New(log, db))
	log.Info("starting server", slog.String("Server Port", cfg.Server_Port))

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"), //The url pointing to API definition
	))

	srv := &http.Server{
		Addr:         cfg.Server_Addr,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	func() error {
		if err := srv.ListenAndServe(); err != nil {
			return fmt.Errorf("%s", err)
		}
		return nil
	}()
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
