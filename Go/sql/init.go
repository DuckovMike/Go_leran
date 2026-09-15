package sql

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func CreateTeamsTable(conn *pgx.Conn, ctx context.Context) {
	sqlStr := `
	CREATE TABLE IF NOT EXISTS teams(
		teamId SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL
	)
	`

	conn.Exec(ctx, sqlStr)
}

func CreateUsersTable(conn *pgx.Conn, ctx context.Context) {
	sqlStr := `
	CREATE TABLE IF NOT EXISTS users(
		userId SERIAL PRIMARY KEY,
		teamId INTEGER NOT NULL REFERECNES teams(teamId) ON DELETE CASCADE,
		fio VARCHAR(100) NOT NULL,
		rate NUMERIC NOT NULL
	);
	`

	conn.Exec(ctx, sqlStr)
}

func CreateAttandanceTable(conn *pgx.Conn, ctx context.Context) {
	sqlStr := `
	CREATE TABLE IF NOT EXISTS attandance(
		id SERIAL PRIMARY KEY,
		userId INTEGER NOT NULL REFERECNES users(userId) ON DELETE CASCADE,
		date DATE NOT NULL,
		attandanceType INT NOT NULL
	);
	`

	conn.Exec(ctx, sqlStr)
}

func ReadTeams(conn *pgx.Conn, ctx context.Context) {
	sqlStr := `
	SELECT * FROM teams;
	`

	conn.Query(ctx, sqlStr)
}

func ReadUsers(conn *pgx.Conn, ctx context.Context) {
	sqlStr := `
	SELECT * FROM users;
	`
	conn.Query(ctx, sqlStr)
}

func ReadAttandance(conn *pgx.Conn, ctx context.Context) {
	sqlStr := `
	SELECT * FROM attandance;
	`
	conn.Query(ctx, sqlStr)
}

func Connect() {
	connStr := "postgress://postgres:0604@localhost:5432/postgres"

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, connStr)

	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	if err = conn.Ping(ctx); err != nil {
		log.Fatal("От Бд нет ответа:", err)
	}

}
