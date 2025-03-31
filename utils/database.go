package utils

import (
	"context"
	"github.com/jackc/pgx/v5"
	"os"
)

func Database() *pgx.Conn {
	url := os.Getenv("DATABASE_URL")

	if con, err := pgx.Connect(context.Background(), url); err != nil {
		panic("Error connecting to database: " + err.Error())
	} else {
		return con
	}

}
