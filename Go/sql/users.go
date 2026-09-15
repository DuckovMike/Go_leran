package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type User struct {
	UserId int
	Team   string
	Fio    string
	Rate   float64
}

func CreateUsersTable(conn *pgx.Conn, ctx context.Context) {
	sqlStr := `
	CREATE TABLE IF NOT EXISTS users(
		userId SERIAL PRIMARY KEY,
		team VARCHAR(100) NOT NULL,
		fio VARCHAR(100) NOT NULL,
		rate NUMERIC NOT NULL
	);
	`

	_, err := conn.Exec(ctx, sqlStr)
	if err != nil {
		panic(err)

	}
}

func CreateUser(conn *pgx.Conn, ctx context.Context, team string, fio string, rate float64) {
	sqlStr := `
	INSERT INTO users(team, fio, rate)
	VALUES ($1, $2, $3);
	`

	_, err := conn.Exec(ctx, sqlStr, team, fio, rate)
	if err != nil {
		panic(err)
	}
}

func ReadUsers(conn *pgx.Conn, ctx context.Context) ([]User, error) {
	var response []User
	sqlStr := `
	SELECT * FROM users;
	`
	rows, err := conn.Query(ctx, sqlStr)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserId, &user.Team, &user.Fio, &user.Rate)
		if err != nil {
			panic(err)
		}
		response = append(response, user)
	}
	return response, err
}

func ReadUserIdByFioAndTeam(conn *pgx.Conn, ctx context.Context, fio string, team string) (int, error) {
	var userId int
	sqlStr := `
	SELECT userId FROM users WHERE fio = $1 AND team = $2;
	`
	err := conn.QueryRow(ctx, sqlStr, fio, team).Scan(&userId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func ReadUsersByTeam(conn *pgx.Conn, ctx context.Context, team string) ([]User, error) {
	var response []User
	sqlStr := `
	SELECT * FROM users WHERE team = $1;
	`
	rows, err := conn.Query(ctx, sqlStr, team)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserId, &user.Team, &user.Fio, &user.Rate)
		if err != nil {
			panic(err)
		}
		response = append(response, user)
	}
	return response, err
}

func UpdateUser(conn *pgx.Conn,
	ctx context.Context,
	userId int,
	fio string,
	rate float64) error {
	sqlStr := `
	UPDATE users
	SET fio = $1, rate = $2
	WHERE userId = $3;
	`
	_, err := conn.Exec(ctx, sqlStr, fio, rate, userId)
	return err
}

func DeleteUser(conn *pgx.Conn, ctx context.Context, userId int) error {
	sqlStr := `
	DELETE FROM users
	WHERE userId = $1;
	`
	_, err := conn.Exec(ctx, sqlStr, userId)
	return err
}
