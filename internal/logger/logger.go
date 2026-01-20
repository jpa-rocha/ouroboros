// Package logger sets up the default level of slog and makes it print color
package logger

import (
	"log/slog"
	"os"
	"time"

	config "ouroboros/internal/config"

	"github.com/lmittmann/tint"
)

// Init initializes the structured logger with configured log level and colored output.
// Reads log level from configuration (INFO, DEBUG, ERROR, WARN), defaults to DEBUG.
// Writes logs to stderr using tint handler for colored, human-readable output.
// Sets the default slog logger globally for use throughout the application.
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
