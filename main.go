package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"todo/api/routes"
	"todo/cmd/cli"
	"todo/config"
	"todo/services"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func Serve(cfg *config.Config) error {
	fmt.Println("API is listening...")

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      routes.Routes(),
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return srv.ListenAndServe()
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
)

func main() {
	cfg := config.MustLoad()

	logger := setupLogger(cfg.Env)
	_ = logger

	services.InitConnection(fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password))

	cfg.Address = "localhost:8083"
	fmt.Println(cfg.Address)

	err := Serve(cfg)
	if err != nil {
		logger.Error("Error", err)
	}
	cli.Execute()
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return logger
}
