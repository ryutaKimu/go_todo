package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type txKey struct{}

var txContextKey txKey = txKey{}

type Postgres struct {
	DB *sql.DB
}

func NewPostgres() *Postgres {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return &Postgres{DB: db}
}

func (p *Postgres) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	if ctx.Value(txContextKey) != nil {
		return fn(ctx)
	}

	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	ctx = context.WithValue(ctx, txContextKey, tx)

	if err := fn(ctx); err != nil {
		if e := tx.Rollback(); e != nil {
			return e
		}
		return err
	}

	return tx.Commit()
}
