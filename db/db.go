package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDatabase() (c *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
