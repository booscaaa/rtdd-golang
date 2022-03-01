package accountrepository

import (
	"context"
	"time"

	"github.com/booscaaa/go-paginate/paginate"
	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
)

func (db accountRepository) Fetch(
	search string,
	descending []string,
	page int,
	itemsPerPage int,
	sort []string,
	status int,
) (*domain.AccountPaginationResponse, error) {
	accounts := []*domain.Account{}
	count := int32(0)

	ctx := context.Background()

	query, queryCount, err := paginate.
		Paginate(`
			SELECT * FROM account WHERE id > 1
		`).
		Desc(descending).
		Sort(sort).
		RowsPerPage(itemsPerPage).
		Page(page).
		SearchBy(search).
		Query()

	if err != nil {
		return nil, err
	}

	{
		rows, err := db.databaseConnection.Query(
			ctx,
			*query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var idA int32
			var loginA string
			var nameA string
			var passwordA string
			var emailA string
			var accountGroupIDA int32
			var createdDateA time.Time
			var activeA bool

			err := rows.Scan(
				&idA,
				&loginA,
				&emailA,
				&nameA,
				&passwordA,
				&createdDateA,
				&accountGroupIDA,
				&activeA,
			)

			if err != nil {
				return nil, err
			}

			account, err := domain.NewAccount(
				idA,
				loginA,
				nameA,
				passwordA,
				emailA,
				createdDateA,
				accountGroupIDA,
				activeA,
			)

			if err != nil {
				return nil, err
			}

			accounts = append(accounts, account)
		}
	}

	{
		db.databaseConnection.QueryRow(
			ctx,
			*queryCount,
		).Scan(&count)
	}

	return &domain.AccountPaginationResponse{
		Items: accounts,
		Total: count,
	}, nil
}
