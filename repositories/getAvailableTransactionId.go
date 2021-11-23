package repositories

import (
	"context"
	"math/rand"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func getAvailableTransactionID() (string, error) {
	var id string
	var available = false
	var err error

	for available != true {
		id = generateTransactionID()
		available, err = checkTransactionID(id)
		if err != nil {
			return "", err
		}
	}

	return id, nil
}

func generateTransactionID() string {
	var possibleNumbers = []rune("0123456789")
	b := make([]rune, 12)
	for i := range b {
		b[i] = possibleNumbers[rand.Intn(len(possibleNumbers))]
	}
	return string(b)
}

func checkTransactionID(id string) (bool, error) {

	var connectionStr = os.Getenv("PostgresConnectionMWallet")

	pool, err := pgxpool.Connect(context.Background(), connectionStr)
	if err != nil {
		return false, err
	}
	defer pool.Close()

	var sqlQuery = `SELECT * FROM otp_history_pay WHERE trans_id=$1`

	resp, err := pool.Exec(context.Background(), sqlQuery, id)
	if err != nil {
		return false, err
	}

	if resp.RowsAffected() == 0 {
		return true, nil
	}
	return false, nil
}
