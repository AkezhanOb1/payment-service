package repositories

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetWalletFromSiteID(sID string, ctx context.Context) (int64, error) {
	purse, err := getPurseFromMWallet(sID, ctx)
	if err != nil {
		return 0, err
	}

	walletID, err := getWalletIDFromEWallet(purse, ctx)
	if err != nil {
		return 0, err
	}

	return walletID, nil
}

func getPurseFromMWallet(sID string, ctx context.Context) (string, error) {

	var connectionStr = os.Getenv("PostgresConnectionMWallet")

	pool, err := pgxpool.Connect(ctx, connectionStr)
	if err != nil {
		return "", err
	}
	defer pool.Close()

	sqlQuery := `SELECT purse FROM otp_sites_payments WHERE site_id=$1;`

	row := pool.QueryRow(ctx, sqlQuery, sID)

	var purse string

	err = row.Scan(
		&purse,
	)

	if err != nil {
		return "", err
	}

	return purse, nil
}

func getWalletIDFromEWallet(purse string, ctx context.Context) (int64, error) {
	var connectionStr = os.Getenv("PostgresConnectionEWallet")

	pool, err := pgxpool.Connect(ctx, connectionStr)
	if err != nil {
		return 0, err
	}
	defer pool.Close()

	sqlQuery := `select id from wallet where number=$1;`

	row := pool.QueryRow(ctx, sqlQuery, purse)

	var walletID int64

	err = row.Scan(
		&walletID,
	)

	if err != nil {
		return 0, err
	}

	return walletID, nil
}
