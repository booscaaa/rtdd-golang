package accountusecase

import (
	"fmt"
	"log"
	"time"

	"github.com/booscaaa/rtdd-golang/microservices/authenticator/core/domain"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func (session accountUsecase) RefreshAuthentication(
	refreshToken string,
	typeToken string,
) (*domain.LoginAccountResponse, error) {
	auth, err := session.repository.GetAuthByRefreshTypeToken(refreshToken, typeToken)
	if err != nil {
		return nil, err
	}

	if auth.Revoked {
		return nil, fmt.Errorf("This token is revoked!")
	}

	account, err := session.repository.GetByID(auth.AccountID)
	if err != nil {
		return nil, err
	}

	if !account.Active {
		return nil, fmt.Errorf("Your account is inative")
	}

	fiveMinutes := time.Now().Add(time.Minute * 5).Unix()
	jwtToken, err := domain.NewJwtToken(*account, fiveMinutes)
	if err != nil {
		log.Fatalln(err)
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwtToken)
	tokenString, err := token.SignedString([]byte(viper.GetString(`hash.bcrypt`)))
	if err != nil {
		log.Fatalln(err)
	}

	jwtAuthToken := domain.LoginAccountResponse{
		Token: tokenString, Refresh: auth.Hash, Type: "refreshToken",
	}

	return &jwtAuthToken, nil
}
