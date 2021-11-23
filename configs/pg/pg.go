package pg

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

// Conn is a
var Conn *pgx.Conn

func init() {
	var err error
	var connectionStr = os.Getenv("PostgresConnectionMWallet")

	Conn, err = pgx.Connect(context.Background(), connectionStr)
	if err != nil {
		log.Fatalln(err)
	}
}
