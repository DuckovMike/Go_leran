package main

import (
	"Go/sql"
	"context"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Connect(ctx)
	if err != nil {
		panic(err)
	}

	/* 	sql.CreateUsersTable(conn, ctx)
	   	sql.CreateAttandanceTable(conn, ctx)
	   	sql.CreateUser(conn, ctx, "Team B", "John Doe", 1.0)
	   	sql.CreateUser(conn, ctx, "Team A", "John Smith", 0.5)
	   	sql.CreateUser(conn, ctx, "Team A", "John Do", 0.9)
		 	resp, _ := sql.ReadUsersByTeam(conn, ctx, "Team A")

	   	fmt.Println("Users in Team A:")

	   	for i, user := range resp {
	   		fmt.Printf("%d - %s (Rate: %.1f)\n", i, user.Fio, user.Rate)
	   	}*/

	sql.CreateAttandanceByUserTeamAndFio(conn, ctx, "Team A", "John Doe", "2026-01-02", 2)

	defer conn.Close(ctx)
}
