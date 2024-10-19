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

func InsertData(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO playing_with_neon(name, value) SELECT LEFT(md5(i::TEXT), 10), random() FROM generate_series(1, 10) s(i);")
	if err != nil {
		panic(err)
	}
	return nil
}

func QueryData(conn *pgx.Conn) ([]map[string]interface{}, error) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM playing_with_neon")
	if err != nil {
		panic(err)
	}
	var data []map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		var value float64
		err := rows.Scan(&id, &name, &value)
		if err != nil {
			panic(err)
		}
		data = append(data, map[string]interface{}{
			"id":    id,
			"name":  name,
			"value": value,
		})
	}
	return data, nil
}
