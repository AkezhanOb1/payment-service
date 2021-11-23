package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"os"
)

func GetUserAccountId(phoneNumber string, ctx context.Context) (string, error) {
	var connectionStr = os.Getenv("PostgresConnectionMWallet")

	pool, err := pgxpool.Connect(ctx, connectionStr)
	if err != nil {
		return "", nil
	}
	defer pool.Close()

	sqlQuery := `SELECT account_id FROM otp_users_accounts WHERE mobile=$1;`

	row := pool.QueryRow(ctx, sqlQuery, phoneNumber)

	var accountID string

	err = row.Scan(
		&accountID,
	)

	if err != nil {
		return "", err
	}

	return accountID, nil
}
