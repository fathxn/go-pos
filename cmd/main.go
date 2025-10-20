package main

import (
	"go-pos/internal/config"
	"go-pos/internal/server"
	"log/slog"
	"os"
)

func main() {
	err := config.LoadConfig("./config.yml")
	if err != nil {
		panic(err)
	}

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	log := slog.New(logHandler)
	slog.SetDefault(log)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
