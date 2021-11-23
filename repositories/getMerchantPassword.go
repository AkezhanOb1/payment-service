package repositories

import (
	"context"

	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

//GetMerchantPasswordRepository is a
func GetMerchantPasswordRepository(sID string, ctx context.Context) (string, error) {
	var connectionStr = os.Getenv("PostgresConnectionMWallet")

	pool, err := pgxpool.Connect(ctx, connectionStr)
	if err != nil {
		return "", err
	}
	defer pool.Close()

	sqlQuery := `SELECT site_passw FROM otp_sites WHERE site_id=$1;`

	row := pool.QueryRow(ctx, sqlQuery, sID)

	var apiPassword string

	err = row.Scan(
		&apiPassword,
	)

	if err != nil {
		return "", err
	}

	return apiPassword, nil
}
