package repositories

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

//GetMerchantSiteNameRepository is a
func GetMerchantSiteNameRepository(sID string, ctx context.Context) (string, error) {

	var connectionStr = os.Getenv("PostgresConnectionMWallet")

	pool, err := pgxpool.Connect(ctx, connectionStr)
	if err != nil {
		return "", err
	}
	defer pool.Close()

	sqlQuery := `SELECT site_name FROM otp_sites WHERE site_id=$1;`

	row := pool.QueryRow(ctx, sqlQuery, sID)

	var siteName string

	err = row.Scan(
		&siteName,
	)

	if err != nil {
		return "", err
	}

	return siteName, nil
}
