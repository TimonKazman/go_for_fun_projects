package helper

import (
	"log/slog"
	"os"
)

//NOTE: not implemented

type LogHandler struct {
	logger *slog.Logger
	file   *os.File
}

func InitLogHandler() (*LogHandler, error) {
	file, err := os.OpenFile("debug_logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	jsonHandler := slog.NewJSONHandler(file, nil)
	logger := slog.New(jsonHandler)

	return &LogHandler{
		logger: logger,
		file:   file,
	}, nil
}

// Close closes the log file. Should be called to release resources.
func (lh *LogHandler) Close() error {
	if lh.file != nil {
		return lh.file.Close()
	}
	return nil
}
