package logger

import (
	"log/slog"
	"os"
	"strings"
)

var Log *slog.Logger

func Init(levelStr string) {
	level := new(slog.LevelVar)

	switch strings.ToLower(levelStr) {
	case "debug":
		level.Set(slog.LevelDebug)
	case "info":
		level.Set(slog.LevelInfo)
	case "warn", "warning":
		level.Set(slog.LevelWarn)
	case "error":
		level.Set(slog.LevelError)
	default:
		level.Set(slog.LevelInfo)
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	Log = slog.New(handler)
}
