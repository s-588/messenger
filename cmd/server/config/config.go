package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/s-588/messenger/cmd/server/database"
	"github.com/spf13/viper"
)

type ServerCfg struct {
	Port      string `mapstructure:"PORT"`
	Log_path  string `mapstructure:"LOG_FILE"` // path to log file
	Log_level string `mapstructure:"LOG_LEVEL"`
}

func LoadConfigs() (*ServerCfg, *database.PostgresCfg, error) {
	if err := godotenv.Load("server.env", "postgres.env"); err != nil {
		return nil, nil, fmt.Errorf("can't load environment variable: %w", err)
	}

	if err := setServerDefaults(); err != nil {
		return nil, nil, fmt.Errorf("can't load server defaults: %w", err)
	}
	setDatabaseDefaults()
	viper.AutomaticEnv()

	if err := checkEssentialEnvs(); err != nil {
		return nil, nil, fmt.Errorf("essential environment variables is not set: %w", err)
	}

	cfg := &ServerCfg{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("can't unmarshal server config: %w", err)
	}

	pgCfg := &database.PostgresCfg{}
	err = viper.Unmarshal(pgCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("can't unmarshal postgres config: %w", err)
	}

	return cfg, pgCfg, nil
}

// setDatabaseDefaults set default values for Postgres config.
func setDatabaseDefaults() {
	viper.SetDefault("PG_HOST", "localhost")
	viper.SetDefault("PG_PORT", "5432")
	viper.SetDefault("PG_SSL", "disable")
}

// setServerDefaults set default values for server config.
func setServerDefaults() error {
	viper.SetDefault("PORT", "8080")
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("can't retrive home directory of user: %w", err)
	}
	logPath := filepath.Join(home, ".local", "share", "messenger")
	if err = os.MkdirAll(logPath, 0755); err != nil {
		return fmt.Errorf("can't create local folder: %w", err)
	}
	viper.SetDefault("LOG_FILE", logPath)
	viper.SetDefault("LOG_LEVEL", "INFO")
	return nil
}

// checkEssentialEnvs check if essential for work values is not defined in config.
func checkEssentialEnvs() error {
	if viper.GetString("PG_USER") == "" {
		return errors.New("Postgres user is not set. " +
			"Set it with PG_USER parameter in .env file.")
	}
	if viper.GetString("PG_PASSWORD") == "" {
		return errors.New("Postgres password is not set. " +
			"Set it with PG_PASSWORD parameter in .env file.")
	}
	if viper.GetString("PG_DBNAME") == "" {
		return errors.New("Postgres database is not set. " +
			"Set it with PG_DBNAME parameter in .env file.")
	}
	return nil
}
