package middleware

import (
	"net/http"
	"stoneBanking/app/common/utils"
	customError "stoneBanking/app/domain/errors"
)

func GetToken(r *http.Request) (string, error) {
	headerToken := r.Header.Get("Authorization")
	if headerToken == "" {
		return "", customError.ErrorServerTokenNotFound
	}

	tokenID, err := utils.ExtractClaims(headerToken)
	if err != nil {
		return "", err
	}
	return tokenID, nil
}
