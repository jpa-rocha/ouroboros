// Package logger sets up the default level of slog and makes it print color
package logger

import (
	"log/slog"
	"os"
	"time"

	config "ouroboros/internal/config"

	"github.com/lmittmann/tint"
)

func Init() {
	w := os.Stderr

	opts := &slog.HandlerOptions{}

	// Setting loglevel.
	switch config.Opt.Logger.Level {
	case "INFO":
		opts.Level = slog.LevelInfo
	case "DEBUG":
		opts.Level = slog.LevelDebug
	case "ERROR":
		opts.Level = slog.LevelError
	case "WARN":
		opts.Level = slog.LevelWarn
	default:
		opts.Level = slog.LevelDebug
	}

	slog.SetDefault(
		slog.New(tint.NewHandler(w, &tint.Options{Level: opts.Level, TimeFormat: time.Kitchen})),
	)
}
