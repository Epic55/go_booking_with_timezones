package app

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func (a *App) initDBConn() error {
	databaseURL := "postgres://postgres:1@localhost:5432/postgres?sslmode=disable" //os.Getenv("DATABASE_URL")

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return err
	}

	config.MaxConns = 10
	config.MaxConnIdleTime = 30 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return err
	}
	fmt.Println("------db - ", db)
	a.closers = append(a.closers, func() error {
		db.Close()
		return nil
	})

	a.dbConn = db

	return nil
}
