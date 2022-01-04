package account

import (
	"context"
	"os"
	"stoneBanking/app/application/validations"
	"stoneBanking/app/application/vo/input"
	"stoneBanking/app/application/vo/output"
	"stoneBanking/app/common/utils"
	"stoneBanking/app/domain/entities/account"
	"time"

	"github.com/golang-jwt/jwt"
)

type ClaimStruct struct {
	jwt.StandardClaims
	User_id int
}

func (usecase *usecase) LoginUser(ctx context.Context, loginInput input.LoginVO) (output.LoginOutputVO, error) {
	loginInput.CPF = utils.TrimCPF(loginInput.CPF)
	_, err := validations.ValidateLoginInputData(loginInput)
	if err != nil {
		return output.LoginOutputVO{}, err
	}

	account, err := usecase.accountRepository.GetByCPF(context.Background(), loginInput.CPF)
	if err != nil {
		return output.LoginOutputVO{}, errorAccountLogin
	}

	if !input.ValidateHash(account.Secret, loginInput.Secret) {
		return output.LoginOutputVO{}, errorAccountLogin
	}

	token, err := generetateToken(account)
	if err != nil {
		return output.LoginOutputVO{}, errorAccountTokenGeneration
	}

	return output.LoginOutputVO{Token: token}, nil
}

func generetateToken(accountData account.Account) (signedToken string, err error) {
	mySigningKey := os.Getenv("SIGN_KEY")
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
