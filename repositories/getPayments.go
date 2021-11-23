package repositories

import (
	"context"
	"log"
	"time"

	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/AkezhanOb1/payment/models"
)

func GetPaymentsRepository(walletID int64, filter models.ParticipantsFilter, ctx context.Context) (*models.ParticipantsList, error) {

	var connectionStr = os.Getenv("PostgresConnectionEWallet")

	pool, err := pgxpool.Connect(ctx, connectionStr)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	sqlQuery := `
	select
       public.transaction.id,
       public.user.id,
       w.id,
       public.user.phone,
       public.transaction.amount,
	   public.transaction.created_at,
       max(case when kyc_template_items.id = 52 then user_profile.value end) as iin,
	   max(case when kyc_template_items.id = 53 then user_profile.value end) as name,
	   max(case when kyc_template_items.id = 54 then user_profile.value end) as surname,
	   max(case when kyc_template_items.id = 55 then user_profile.value end) as patronymic
    from public.transaction
        inner join wallet w
                on transaction.wallet_from_id = w.id
        inner join global_wallet_owner gwo
                on w.global_wallet_owner_id = gwo.id
        inner join public.user
                on gwo.user_id = public.user.id
        left join public.user_profile
                on public.user.id = user_profile.user_id
        left join kyc_template_items
                on user_profile.kyc_template_item_id = kyc_template_items.id
    where wallet_to_id = $1
        and transaction_type_id=2
        and gwo.user_id is not null
        and public.transaction.created_at >= $2
        and public.transaction.created_at <= $3
		and public.transaction.amount >= $4
        and public.transaction.amount <= $5
    group by public.transaction.id,
             public.user.id,
             gwo.user_id,
             w.id,
             public.transaction.wallet_from_id,
             public.transaction.wallet_to_id, public.transaction.amount
    order by public.transaction.id, public.transaction.created_at;`

	rows, err := pool.Query(
		ctx,
		sqlQuery,
		walletID,
		filter.Period.FromDate,
		filter.Period.ToDate,
		//filter.Sum.FromSum,
		//filter.Sum.ToSum,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var participants models.ParticipantsList

	for rows.Next() {
		var participant models.Participant
		var t time.Time
		err = rows.Scan(
			&participant.TransactionID,
			&participant.UserID,
			&participant.WalletID,
			&participant.PhoneNumber,
			&participant.Amount,
			&t,
			&participant.IIN,
			&participant.FirstName,
			&participant.SecondName,
			&participant.Patronymic,
		)

		participant.CreatedAt = t.Format(time.RFC3339)
		participants.Participants = append(participants.Participants, participant)
	}

	if err != nil {
		return nil, err
	}
	log.Println(len(participants.Participants))

	return &participants, nil
}
