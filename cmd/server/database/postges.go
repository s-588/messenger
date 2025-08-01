package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/s-588/messenger/cmd/server/database/postgres"
)

type Postgres struct {
	querys *postgres.Queries
}

type PostgresCfg struct {
	Username string `mapstructure:"DB_USER"`
	Password string `mapstructure:"PG_PASSWORD"`
	Host     string `mapstructure:"PG_HOST"`
	Port     string `mapstructure:"PG_PORT"`
	Dbname   string `mapstructure:"PG_DBNAME"`
	Ssl      string `mapstructure:"PG_SSL"`
}

func NewPostgres(ctx context.Context, cfg *PostgresCfg) (*Postgres, error) {
	connstr := GetConnectionString(cfg)
	conn, err := pgxpool.New(ctx, connstr)
	if err != nil {
		return nil, err
	}
	return &Postgres{
		querys: postgres.New(conn),
	}, nil
}

func GetConnectionString(cfg *PostgresCfg) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname, cfg.Ssl,
	)
}
