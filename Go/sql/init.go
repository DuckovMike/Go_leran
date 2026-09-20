package sql

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context) (*pgx.Conn, error) {
	connStr := "postgres://postgres:0604@localhost:5432/postgres"

	conn, err := pgx.Connect(ctx, connStr)

	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	if err = conn.Ping(ctx); err != nil {
		log.Fatal("От Бд нет ответа:", err)
	}
	return conn, err
}

func CreateBaseAttandance(conn *pgx.Conn,
	ctx context.Context,
	startDate string,
	endDate string) {

	sqlStr := `
	CREATE TABLE IF NOT EXISTS baseAttandance(
		id SERIAL PRIMARY KEY,
		date DATE NOT NULL UNIQUE,
		attandanceType INT NOT NULL,
		attandanceHours INT
	);
	`
	_, err := conn.Exec(ctx, sqlStr)
	if err != nil {
		panic(err)
	}

	sqlStr = `
	INSERT INTO baseAttandance(date, attandanceType)
	SELECT 
		d,
		CASE WHEN EXTRACT(DOW FROM d) IN (0, 6) THEN 1 ELSE 0 END
	FROM generate_series($1::date, $2::date, '1 day'::interval) AS d
	`
	_, err = conn.Exec(ctx, sqlStr, startDate, endDate)
	if err != nil {
		panic(err)
	}
}
