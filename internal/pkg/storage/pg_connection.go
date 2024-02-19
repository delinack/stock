package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// PGConfig struct for postgres config
type PGConfig struct {
	User           string `mapstructure:"POSTGRES_USER" validate:"required"`
	Password       string `mapstructure:"POSTGRES_PASSWORD" validate:"required"`
	Host           string `mapstructure:"POSTGRES_HOST" validate:"required,hostname_rfc1123"`
	Port           uint64 `mapstructure:"POSTGRES_PORT" validate:"required,gte=0,lte=65535"`
	DBName         string `mapstructure:"POSTGRES_DB" validate:"required"`
	SimpleProtocol bool   `mapstructure:"SIMPLE_PROTOCOL" validate:"required"`
}

// NewPGConnection - Connection parameters are validate inside.
// You don't need to import database library driver.
// Support dsn parameters
func NewPGConnection(ctx context.Context, cfg PGConfig) (*sqlx.DB, error) {
	var confDsn pgx.ConnConfig
	var err error

	confPgx := pgx.ConnConfig{
		Host:                 cfg.Host,
		Port:                 uint16(cfg.Port),
		User:                 cfg.User,
		Password:             cfg.Password,
		Database:             cfg.DBName,
		PreferSimpleProtocol: cfg.SimpleProtocol,
	}

	//Add dsn config to config from incoming parameters, they not overwrite
	conf := confDsn.Merge(confPgx)

	//Wrapped *sql.DB building on pgx driver and incoming configuration
	db := sqlx.NewDb(stdlib.OpenDB(conf), "pgx")

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("cannot ping connection: %w", err)
	}

	return db, nil
}
