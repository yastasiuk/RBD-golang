package main

import (
	"log/slog"
	"os"
)

func setupLogger() *slog.Logger {
	// Console handler (e.g., os.Stdout)
	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	// Wrap them with LevelBasedHandler
	logger := slog.New(consoleHandler)

	slog.SetDefault(logger)
	return logger
}
