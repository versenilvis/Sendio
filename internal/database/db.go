// Package database - db (PostgreSQL)
package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/versenilvis/sendio/internal/config"
	"github.com/versenilvis/sendio/internal/logger"
)

type DB struct {
	Pool *pgxpool.Pool
}

// NewConnection creates a new database connection pool
// parse config: convert db url to a struct that pgx understands
// create pool but still not connected yet
// ping to check
func NewConnection(cfg *config.Config) (*DB, error) {
	if cfg.Database.URL == "" {
		// please dont use panic here, just print log and return error
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	poolConfig, err := pgxpool.ParseConfig(cfg.Database.URL)
	if err != nil {
		return nil, fmt.Errorf("error parsing database config: %w", err)
	}

	// NOTE: change in env

	// maximum number of connections allowed in the pool
	poolConfig.MaxConns = int32(cfg.Database.MaxConns)
	// minimum number of connections to maintain in the pool even if idle
	poolConfig.MinConns = int32(cfg.Database.MinConns)
	// duration after which an idle connection will be closed and removed from the pool
	poolConfig.MaxConnIdleTime = 5 * time.Minute
	// maximum amount of time a connection may be reused
	poolConfig.MaxConnLifetime = 60 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("error creating database pool: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	logger.Log.Info("Connected to PostgreSQL successfully")

	return &DB{Pool: pool}, nil
}

// Close closes db
// take note that this is not the same as utils.Close function
func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}

// RunMigrations runs goose migrations
// because pgx is designed for best performance for postgresql,
// it doesnt follow the rules of go's database/sql interface
// using stdlib of pgx to create an adapter, take the pool connection configuration
// then wrap it to be a *sql.DB that goose can understand
func (db *DB) RunMigrations() error {
	// goose needs a *sql.DB
	sqlDB := stdlib.OpenDB(*db.Pool.Config().ConnConfig)
	defer sqlDB.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	if err := goose.Up(sqlDB, "migrations"); err != nil {
		return fmt.Errorf("goose migration failed: %w", err)
	}

	logger.Log.Info("Migrations completed successfully with goose")
	return nil
}
