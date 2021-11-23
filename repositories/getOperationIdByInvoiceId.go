package repositories

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetOperationIdRepository(invoiceID string, ctx context.Context) (string, error) {
	var connectionStr = os.Getenv("PostgresConnectionMWallet")

	pool, err := pgxpool.Connect(ctx, connectionStr)
	if err != nil {
		return "", nil
	}
	defer pool.Close()

	sqlQuery := `select payment_id_pay from otp_history_pay where invoice_id=$1 and status_pay=3;`

	row := pool.QueryRow(ctx, sqlQuery, invoiceID)

	var operationID string

	err = row.Scan(
		&operationID,
	)

	if err != nil {
		return "", err
	}

	return operationID, nil
}
