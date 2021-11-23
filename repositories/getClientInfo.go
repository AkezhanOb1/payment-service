package repositories

import (
	"context"

	"os"

	"github.com/AkezhanOb1/payment/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

//GetClientInfoRepository is
func GetClientInfoRepository(phoneNumber string, ctx context.Context) (*models.User, error) {
	var connectionStr = os.Getenv("PostgresConnectionEWallet")

	pool, err := pgxpool.Connect(ctx, connectionStr)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	sqlQuery := `
		select
		   public.user.id,
		   public.user.phone,
		   max(case when kyc_template_items.id = 52 then user_profile.value end) as iin,
		   max(case when kyc_template_items.id = 53 then user_profile.value end) as name,
		   max(case when kyc_template_items.id = 54 then user_profile.value end) as surname,
		   max(case when kyc_template_items.id = 55 then user_profile.value end) as patronymic,
		   w.id
		from public.user
			left join public.user_profile
					on public.user.id = user_profile.user_id
			left join kyc_template_items
					on user_profile.kyc_template_item_id = kyc_template_items.id
			left join global_wallet_owner gwo on public.user.id = gwo.user_id
			inner join wallet w on gwo.id = w.global_wallet_owner_id
			where public.user.phone = $1
			and public.user.is_active = true
			and w.type_id = 1
		group by
				 public.user.id,
				 w.id
		order by public.user.id;`

	var user models.User

	row := pool.QueryRow(ctx, sqlQuery, phoneNumber)

	err = row.Scan(
		&user.ID,
		&user.PhoneNumber,
		&user.IIN,
		&user.FirstName,
		&user.LastName,
		&user.Patronymic,
		&user.WalletID,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
