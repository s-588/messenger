package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

type LoggerConf struct {
	LogPath   string
	LogLevel string
}

func SetUpSlog(config LoggerConf) error {
	logFile, err := os.OpenFile(config.LogPath, os.O_CREATE&os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("can't create log file: %w", err)
	}
	w := io.MultiWriter(logFile, os.Stdout)
	handler := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: getDebugLevel(config.LogLevel),
	}))
	slog.SetDefault(handler)
	return nil
}

func getDebugLevel(level string) slog.Leveler {
	switch strings.ToTitle(level) {
	case "INFO":
		return slog.LevelInfo
	case "DEBUG":
		return slog.LevelDebug
	case "WARN":
		return slog.LevelWarn
	}
	return slog.LevelError
}
