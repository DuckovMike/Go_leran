package sql

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func test() {
	connStr := "postgress://postgres:0604@localhost:5432/postgres"

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	if err = conn.Ping(ctx); err != nil {
		log.Fatal("От Бд нет ответа:", err)
	}

	sqlStr := `
	CREATE TABLE IF NOT EXISTS test1(
		id SERIAL PRIMARY KEY,
		fio VARCHAR(50) NOT NULL,
	);
	`

	conn.Exec(ctx, sqlStr)
}
