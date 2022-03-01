package domain

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountUsecase interface {
	Create(account *Account) (*Account, error)
	Update(account *Account) (*Account, error)
	Delete(id int32) (*Account, error)
	GetByID(id int32) (*Account, error)
	GetByName(name string) (*Account, error)
	Fetch(
		account Account,
		search string,
		descending []string,
		page int,
		itemsPerPage int,
		sort []string,
		status int,
	) (*AccountPaginationResponse, error)
	// Authentication
	VerifyAuthentication(login string, password string) (*LoginAccountResponse, error)
	RefreshAuthentication(refreshToken string, typeToken string) (*LoginAccountResponse, error)
}

type AccountRepository interface {
	Create(
		login string,
		name string,
		password string,
		email string,
		accountGroupID int32,
		active bool,
	) (*Account, error)
	Update(
		id int32,
		login string,
		name string,
		email string,
		accountGroupID int32,
		active bool,
	) (*Account, error)
	UpdateWithPassword(
		id int32,
		login string,
		name string,
		password string,
		email string,
		accountGroupID int32,
		active bool,
	) (*Account, error)
	Delete(id int32) (*Account, error)
	GetByID(id int32) (*Account, error)
	GetByName(name string) (*Account, error)
	GetByEmail(email string) (*Account, error)
	Fetch(
		search string,
		descending []string,
		page int,
		itemsPerPage int,
		sort []string,
		status int,
	) (*AccountPaginationResponse, error)

	//Authentication
	GetAccountByLoginPassword(login string, email string, password string) (*Account, error)
	GetAuthByRefreshTypeToken(refreshToken string, typeToken string) (*Auth, error)
	GetAuthByAccountID(id int32) (*Auth, error)
	CreateAuth(
		typeToken string,
		hash string,
		token string,
		accountID int32,
		revoked bool,
	) (*Auth, error)
	DeleteAuth(id int32) (*Auth, error)
}

func (account Account) isValidFromStruct() error {
	if account.Id == 0 {
		return fmt.Errorf("The account ID not be empty")
	}

	if account.Login == "" {
		return fmt.Errorf("The account LOGIN not be empty")
	}

	if account.Name == "" {
		return fmt.Errorf("The account NAME not be empty")
	}

	if account.Password == "" {
		return fmt.Errorf("The account PASSWORD not be empty")
	}

	if account.Email == "" {
		return fmt.Errorf("The account EMAIL not be empty")
	}

	if account.AccountGroupID == 0 {
		return fmt.Errorf("The account ACCOUNT GROUP ID not be empty")
	}

	return nil
}

func NewAccount(
	id int32,
	login string,
	name string,
	password string,
	email string,
	createdDate time.Time,
	accountGroupID int32,
	active bool,
) (*Account, error) {
	account := Account{
		Id:             id,
		Login:          login,
		Name:           name,
		Password:       password,
		Email:          email,
		CreatedDate:    timestamppb.New(createdDate),
		AccountGroupID: accountGroupID,
		Active:         active,
	}

	if err := account.isValidFromStruct(); err != nil {
		return nil, err
	}

	return &account, nil
}
