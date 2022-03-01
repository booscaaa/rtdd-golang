package accountusecase

import "github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"

func (usecase accountUsecase) Fetch(
	account domain.Account,
	search string,
	descending []string,
	page int,
	itemsPerPage int,
	sort []string,
	status int,
) (*domain.AccountPaginationResponse, error) {
	accountPagination, err := usecase.repository.Fetch(
		search,
		descending,
		page,
		itemsPerPage,
		sort,
		status,
	)

	if err != nil {
		return nil, err
	}

	return accountPagination, nil
}
