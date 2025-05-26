package database

import (
	"context"
	"ostkost/go-ps-hw-fiber/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDbPool(config *config.DatabaseConfig) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.DbUrl)
	if err != nil {
		panic(err)
	}
	return dbpool
}
