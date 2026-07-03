package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect(connectionString string) (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), connectionString)
}
