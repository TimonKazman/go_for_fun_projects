package helper

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

func InitializeLogging(debug_mode bool) *slog.Logger {
	var opts *slog.HandlerOptions

	if debug_mode {
		opts = &slog.HandlerOptions{Level: slog.LevelDebug}
	} else {
		opts = nil
	}

	logFile, err := os.OpenFile("debug_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(fmt.Sprintf("opening log file not possible: %v", err))
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	jsonHandler := slog.NewJSONHandler(multiWriter, opts)
	logger := slog.New(jsonHandler)

	return logger
}
