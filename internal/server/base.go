package server

import (
	"go-pos/internal/config"
	"go-pos/internal/infra/database"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Start() error {
	cfg := config.GetConfig()

	db, err := database.ConnectPostgres(cfg.DB)
	if err != nil {
		return err
	}

	_ = db

	router := chi.NewRouter()

	slog.Info("server "+cfg.App.Name+" starting", slog.Any("port", cfg.App.Port))

	http.ListenAndServe(":"+cfg.App.Port, router)

	return nil
}
