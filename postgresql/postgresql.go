package postgresql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func connect() *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected!\n")
	return dbpool
}

func Query(q *string) []string {
	res, _ := connect().Query(context.Background(), *q)
	row, err := pgx.CollectRows(res, pgx.RowTo[string])
	if err != nil {
		fmt.Println(err)
	}
	return row
}

func QueryRow(q *string) (res string) {
	err := connect().QueryRow(context.Background(), *q).Scan(&res)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func Exec(q *string) {
	_, err := connect().Exec(context.Background(), *q)
	if err != nil {
		fmt.Println(err)
	}
}
