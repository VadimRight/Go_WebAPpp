package main

import (
	"fmt"
	"log/slog"
	"os"
	"net/http"
	"github.com/VadimRight/Go_WebApp/internal/config"
	"github.com/VadimRight/Go_WebApp/internal/lib/logger/handlers/slogpretty"
	mwlogger "github.com/VadimRight/Go_WebApp/internal/server/middleware/logger"
	"github.com/VadimRight/Go_WebApp/internal/storage/postgres"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/VadimRight/Go_WebApp/internal/server/handler/url"
	"github.com/swaggo/http-swagger/v2"
	_ "github.com/VadimRight/Go_WebApp/cmd/url-shortener/docs"
)


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
	test_add, err := postgres.TestAddUrl()
	if err != nil {
		log.Error("failed to init storage", err)
		os.Exit(1)
	}
	fmt.Println(test_add)
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
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"),
		httpSwagger.BeforeScript(`const SomePlugin = (system) => ({
    // Some plugin
  });
`),
		httpSwagger.AfterScript(`const someOtherCode = function(){
    // Do something
  };
  someOtherCode();`),
		httpSwagger.Plugins([]string{
			"SomePlugin",
			"AnotherPlugin",
		}),
		httpSwagger.UIConfig(map[string]string{
			"showExtensions":        "true",
			"onComplete":            `() => { window.ui.setBasePath('v3'); }`,
			"defaultModelRendering": `"model"`,
		}),
	))
	srv := &http.Server{
		Addr:		cfg.Server_Addr,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	func () error {
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
