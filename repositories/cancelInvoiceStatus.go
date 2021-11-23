package repositories

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"os"

	"github.com/AkezhanOb1/payment/models"
)

func CancelInvoiceRepository(payment models.Payment, accountID string, ctx context.Context) error {

	var connectionStr = os.Getenv("PostgresConnectionMWallet")

	transactionID, err := getAvailableTransactionID()
	if err != nil {
		return err
	}

	pool, err := pgxpool.Connect(ctx, connectionStr)
	if err != nil {
		return err
	}
	defer pool.Close()

	var sqlQuery = `INSERT INTO otp_history_pay
    					(trans_id, invoice_id, account_id, amount, fee, currency, status_pay, mark, wl)
					VALUES ($1, $2, $3, $4, 0, 'KZT', 2, 1, 'bloomzed');`

	_, err = pool.Exec(ctx, sqlQuery, transactionID, payment.InvoiceID, accountID, payment.Amount.Sum)

	if err != nil {
		return err
	}
	return nil
}
