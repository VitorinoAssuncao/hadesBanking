package middleware

import (
	"net/http"
	"stoneBanking/app/common/utils"
	customError "stoneBanking/app/domain/errors"
)

func GetAccountIDFromToken(r *http.Request) (string, error) {
	headerToken := r.Header.Get("Authorization")
	if headerToken == "" {
		return "", customError.ErrorServerTokenNotFound
	}

	accountExternalID, err := utils.ExtractClaims(headerToken)
	if err != nil {
		return "", err
	}
	return accountExternalID, nil
}
