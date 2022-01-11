package middleware

import (
	"net/http"
	"stoneBanking/app/domain/entities/token"
	customError "stoneBanking/app/domain/errors"
)

func GetAccountIDFromToken(r *http.Request, t token.TokenInterface) (string, error) {
	headerToken := r.Header.Get("Authorization")
	if headerToken == "" {
		return "", customError.ErrorServerTokenNotFound
	}

	accountExternalID, err := t.ExtractAccountIDFromToken(headerToken)
	if err != nil {
		return "", err
	}
	return accountExternalID, nil
}
