package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Attandance struct {
	Id              int
	UserId          int
	Date            string
	AttandanceType  int
	AttandanceHours int
}

func CreateAttandanceTable(conn *pgx.Conn, ctx context.Context) error {
	sqlStr := `
	CREATE TABLE IF NOT EXISTS attandance(
		id SERIAL PRIMARY KEY,
		userId INTEGER NOT NULL REFERENCES users(userId) ON DELETE CASCADE,
		date DATE NOT NULL,
		attandanceType INT NOT NULL,
		attandanceHoyrs INT
	);
	`

	_, err := conn.Exec(ctx, sqlStr)
	return err
}

func UpdateAttandanceByUserT_F_D(
	conn *pgx.Conn,
	ctx context.Context,
	team string,
	fio string,
	date string,
	attandanceType int,
	attandanceHours int) error {

	sqlStr := `
	UPDATE attandance 
	SET attandanceType = $4, attandanceHours = $5
	WHERE userId = (SELECT userId FROM users WHERE team = $1 AND fio = $2) AND date = $3
	`

	_, err := conn.Exec(ctx, sqlStr, team, fio, date, attandanceType, attandanceHours)
	return err
}

func CreateAttandanceByUserT_F(
	conn *pgx.Conn,
	ctx context.Context,
	team string,
	fio string,
	date string,
	attandanceType int,
	attandanceHours int) error {

	sqlStr := `
	INSERT INTO attandance(userId, date, attandaceType, attandanceHours)
	VALUES ((SELECT userId
		FROM users
		WHERE fio = $1 AND team = $2), 
		$3, 
		$4, 
		$5)
	`

	_, err := conn.Exec(ctx, sqlStr, fio, team, date, attandanceType, attandanceHours)
	return err
}

func ReadAttandanceByUserId(conn *pgx.Conn, ctx context.Context, userId int) {
	sqlStr := `
	SELECT * FROM attandance
	WHERE userId = $1;
	`
	conn.Query(ctx, sqlStr, userId)
}

func ReadAttandanceByUserT_F_D(
	conn *pgx.Conn,
	ctx context.Context,
	team string,
	fio string,
	date string) Attandance {
	sqlStr := `
	SELECT id, userId, date, attandanceType, attandanceHours
	FROM attandance
	WHERE userId = (SELECT userId FROM users WHERE team = $1 AND fio = $2) AND date = $3::date
	`
	var result Attandance

	conn.QueryRow(ctx, sqlStr, team, fio, date).Scan(
		&result.Id,
		&result.UserId,
		&result.Date,
		&result.AttandanceType,
		&result.AttandanceHours)

	return result
}

func ReadAttandanceByUserT_F(
	conn *pgx.Conn,
	ctx context.Context,
	team string,
	fio string,
	date string) []Attandance {
	sqlStr := `
	SELECT id, userId, date, attandanceType, attandanceHours
	FROM attandance
	WHERE userId = (SELECT userId FROM users WHERE team = $1 AND fio = $2) AND date = $3::date
	`
	var result []Attandance

	rows, err := conn.Query(ctx, sqlStr, team, fio, date)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var att Attandance
		rows.Scan(&att.Id, &att.UserId, &att.Date, &att.AttandanceType, &att.AttandanceHours)
		result = append(result, att)
	}

	return result
}

func DeleteAttandanceByUserT_F_D(
	conn *pgx.Conn,
	ctx context.Context,
	team string,
	fio string,
	date string) error {

	sqlStr := `
	DELETE FROM attandance
	WHERE 
		userId = (SELECT userId 
					FROM users 
					WHERE team = $1 AND fio = $2) 
		AND date = $3;
	`

	_, err := conn.Exec(ctx, sqlStr, team, fio, date)
	return err
}

func DeleteAttandanceByUserT_F(
	conn *pgx.Conn,
	ctx context.Context,
	team string,
	fio string) error {

	sqlStr := `
	DELETE FROM users
	WHERE userId = (SELECT userId FROM users WHERE team = $1 AND fio = $2)
	`

	_, err := conn.Exec(ctx, sqlStr, team, fio)
	return err
}
