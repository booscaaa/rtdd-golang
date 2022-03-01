package domain

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type JwtToken struct {
	Account Account `json:"account,omitempty"`
	Exp     int64   `json:"exp,omitempty"`
	jwt.StandardClaims
}

func (jwtToken JwtToken) isValidFromStruct() error {
	if jwtToken.Account.Id == 0 {
		return fmt.Errorf("The JWT TOKEN ACCOUNT not be empty")
	}

	if jwtToken.Exp == 0 {
		return fmt.Errorf("The JWT TOKEN EXP not be empty")
	}

	return nil
}

func NewJwtToken(account Account, exp int64) (*JwtToken, error) {
	jwtToken := JwtToken{
		Account: account,
		Exp:     exp,
	}

	if err := jwtToken.isValidFromStruct(); err != nil {
		return nil, err
	}

	return &jwtToken, nil
}
