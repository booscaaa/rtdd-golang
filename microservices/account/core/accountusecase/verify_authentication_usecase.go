package accountusecase

import (
	"fmt"
	"time"

	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func (session accountUsecase) VerifyAuthentication(
	login string,
	password string,
) (*domain.LoginAccountResponse, error) {
	account, err := session.repository.GetAccountByLoginPassword(login, login, password)
	if err != nil {
		return nil, err
	}

	fiveMinutes := time.Now().Add(time.Minute * 5).Unix()
	jwtToken, err := domain.NewJwtToken(*account, fiveMinutes)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwtToken)
	tokenString, err := token.SignedString([]byte(viper.GetString(`hash.bcrypt`)))
	if err != nil {
		return nil, err
	}

	auth, err := session.repository.GetAuthByAccountID(account.Id)
	var hash string

	if err != nil {
		if err.Error() == "Token not created yet" {
			password := []byte(viper.GetString(`hash.bcrypt`))
			hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
			hash = string(hashedPassword)

			_, err := session.repository.CreateAuth(
				"refreshToken",
				hash,
				tokenString,
				account.Id,
				false,
			)

			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	if auth != nil {
		if auth.Revoked {
			return nil, fmt.Errorf("Your account has been revoked")
		}

		hash = auth.Hash
	}

	jwtAuthToken := domain.LoginAccountResponse{
		Token: tokenString, Refresh: hash, Type: "refreshToken",
	}

	return &jwtAuthToken, nil
}
