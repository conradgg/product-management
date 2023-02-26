package postgresql

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func Postgresql() {
	dbUrl := "postgres://test:test@localhost:5432/test"
	conn, err := pgx.Connect(context.Background(), os.Getenv(dbUrl))
}
