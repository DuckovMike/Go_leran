package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateAttandanceTable(conn *pgx.Conn, ctx context.Context) {
	sqlStr := `
	CREATE TABLE IF NOT EXISTS attandance(
		id SERIAL PRIMARY KEY,
		userId INTEGER NOT NULL REFERENCES users(userId) ON DELETE CASCADE,
		date DATE NOT NULL,
		attandanceType INT NOT NULL
	);
	`

	_, err := conn.Exec(ctx, sqlStr)
	if err != nil {
		panic(err)

	}
}

func ReadAttandanceByUserId(conn *pgx.Conn, ctx context.Context, userId int) {
	sqlStr := `
	SELECT * FROM attandance
	WHERE userId = $1;
	`
	conn.Query(ctx, sqlStr, userId)
}

func CreateAttandanceByUserTeamAndFio(
	conn *pgx.Conn,
	ctx context.Context,
	team string,
	fio string,
	date string,
	attandanceType int) error {

	sqlStr := `
	

	INSERT INTO attandance(userId, date, attandaceType)
	VALUES ((SELECT userId
		FROM users
		WHERE fio = $1 AND team = $2), 
		date, 
		attandaceType), $3, $4)
	`

	_, err = conn.Exec(ctx, sqlStr, fio, team, date, attandanceType)
	return err
}
