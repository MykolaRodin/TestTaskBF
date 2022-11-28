package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:secretpassword@127.0.0.1:5432/gousers?sslmode=disable")
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
