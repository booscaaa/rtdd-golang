package domain

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (auth Auth) isValidFromStruct() error {
	if auth.Id == 0 {
		return fmt.Errorf("The auth ID not be empty")
	}

	if auth.Type == "" {
		return fmt.Errorf("The auth TYPE not be empty")
	}

	if auth.Hash == "" {
		return fmt.Errorf("The auth HASH not be empty")
	}

	if auth.Token == "" {
		return fmt.Errorf("The auth TOKEN not be empty")
	}

	if auth.AccountID == 0 {
		return fmt.Errorf("The auth ACCOUNT ID not be empty")
	}

	return nil
}

func NewAuth(
	id int32,
	typeToken string,
	hash string,
	token string,
	accountID int32,
	revoked bool,
	createdDate time.Time,
) (*Auth, error) {
	auth := Auth{
		Id:          id,
		Type:        typeToken,
		Hash:        hash,
		Token:       token,
		AccountID:   accountID,
		Revoked:     revoked,
		CreatedDate: timestamppb.New(createdDate),
	}

	if err := auth.isValidFromStruct(); err != nil {
		return nil, err
	}

	return &auth, nil
}
