package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/s-588/messenger/cmd/server/config"
	"github.com/spf13/viper"
)

func main() {
	cfg, pgcfg, err := config.LoadConfigs()
	if err != nil {
		slog.Error("can't load configs", "error", err)
		os.Exit(1)
	}
	slog.Info("config loaded succesfully", "server config", *cfg, "postgres config", *pgcfg)
	setLogger()
	slog.Info("logger settled up")
	// TODO: connect to database
}

func setLogger() error {
	logFile, err := os.OpenFile(viper.GetString("LOG_FILE"),
		os.O_CREATE&os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("can't create log file: %w", err)
	}
	w := io.MultiWriter(logFile, os.Stdout)
	handler := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: getDebugLevel(),
	}))
	slog.SetDefault(handler)
	return nil
}

func getDebugLevel() slog.Leveler {
	var lvl slog.Level
	switch viper.GetString("LOG_LEVEL") {
	case "INFO":
		lvl = slog.LevelInfo
	case "DEBUG":
		lvl = slog.LevelDebug
	case "ERROR":
		lvl = slog.LevelError
	case "WARN":
		lvl = slog.LevelWarn
	}
	return lvl
}
