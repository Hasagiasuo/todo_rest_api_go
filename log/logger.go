package log 

import (
	"log/slog"
	"os"
)

const (
	localEnv = "local"
	servEnv = "server"
)

func LoggerSetup(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case localEnv:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ Level: slog.LevelDebug }))
	case servEnv:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{ Level: slog.LevelDebug }))
	} 
	return log
}