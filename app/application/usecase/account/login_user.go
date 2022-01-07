package account

import (
	"context"
	"os"
	"stoneBanking/app/common/utils"
	"stoneBanking/app/domain/entities/account"
	customError "stoneBanking/app/domain/errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type ClaimStruct struct {
	jwt.StandardClaims
	User_id int
}

func (usecase *usecase) LoginUser(ctx context.Context, loginInput account.Account) (string, error) {

	tempAccount, err := usecase.accountRepository.GetByCPF(context.Background(), loginInput.CPF)
	if err != nil {
		return "", customError.ErrorAccountLogin
	}

	if !utils.ValidateHash(tempAccount.Secret, loginInput.Secret) {
		return "", customError.ErrorAccountLogin
	}

	token, err := generetateToken(tempAccount)
	if err != nil {
		return "", customError.ErrorAccountTokenGeneration
	}

	return token, nil
}

func generetateToken(accountData account.Account) (signedToken string, err error) {
	mySigningKey := []byte(os.Getenv("SIGN_KEY"))
	claims := &ClaimStruct{
		User_id: accountData.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(2 * time.Hour).Unix(),
			Issuer:    "vitorino",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString(mySigningKey)

	if err != nil {
		println("Erro na geração do token")
	}
	return signedToken, nil
}
