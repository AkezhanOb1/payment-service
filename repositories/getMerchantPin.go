package repositories

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetMerchantPinRepository(sID string, ctx context.Context) (string, error) {

	var connectionStr = os.Getenv("PostgresConnectionMWallet")

	pool, err := pgxpool.Connect(ctx, connectionStr)
	if err != nil {
		return "", err
	}
	defer pool.Close()

	sqlQuery := `SELECT pinok FROM otp_sites_personnel WHERE site_id=$1;`

	row := pool.QueryRow(ctx, sqlQuery, sID)

	var pin string

	err = row.Scan(
		&pin,
	)

	if err != nil {
		return "", err
	}

	return pin, nil
}
