package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Context(context.Background())
	db, err := pgx.Connect(ctx, "postgres://postgres:0604@localhost:5432/postgres")

	if err != nil {
		panic(err)
	}

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("OK")
}
