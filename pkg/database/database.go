package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
)

func InitDatabase() (*pgx.Conn, error) {
	connStr := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CreateTable(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS playing_with_neon(id SERIAL PRIMARY KEY, name TEXT NOT NULL, value REAL);")
	if err != nil {
		return err
	}
	return nil
}
