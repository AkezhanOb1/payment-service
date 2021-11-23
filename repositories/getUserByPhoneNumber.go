package repositories

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/AkezhanOb1/payment/models"
)

func GetUserByPhoneNumberRepository(phoneNumber string, ctx context.Context) (interface{}, error) {

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
           public.user.kyc_status,
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
			and w.type_id = 1
		group by
				 public.user.id,
				 w.id
		order by public.user.id;`

	row := pool.QueryRow(ctx, sqlQuery, phoneNumber)
	var client models.Client
	var kyc models.Reference
	var user models.User

	err = row.Scan(
		&user.ID,
		&user.PhoneNumber,
		&kyc.Code,
		&user.IIN,
		&user.FirstName,
		&user.LastName,
		&user.Patronymic,
		&user.WalletID,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return models.CustomError{
				Code:    200,
				Message: "данный немер не зарегистрирован в системе bloomzed",
			}, nil
		}
		return nil, err
	}

	switch kyc.Code {
	case 0:
		kyc.Description = "не идентифицирован"
	case 1:
		kyc.Description = "частично идентифицирован"
	case 2:
		kyc.Description = "идентифицирован"
	}

	client.KYC = &kyc
	client.User = user

	return &client, nil
}
