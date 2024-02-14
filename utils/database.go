package utils

import (
	"context"

	"github.com/jackc/pgx/v4"
)

var DB *pgx.Conn

func InitializeDB(connectionString string) error {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return err
	}
	DB = conn
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close(context.Background())
	}
}
