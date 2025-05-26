package database

import (
	"context"
	"ostkost/go-ps-fiber/pkg/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func CreateDbPool(config *config.DatabaseConfig, logger *zerolog.Logger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Url)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create db pool")
		panic(err)
	}
	return dbpool
}
