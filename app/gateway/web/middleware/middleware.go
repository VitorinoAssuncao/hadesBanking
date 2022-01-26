package middleware

import (
	"net/http"
	logHelper "stoneBanking/app/common/utils/logger"
	"stoneBanking/app/domain/entities/token"
	customError "stoneBanking/app/domain/errors"
)

func GetAccountIDFromToken(r *http.Request, t token.Repository) (string, error) {

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

func LogRoutes(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logHelper.NewLogger()
		log.LogInfo("request", "received request in url: "+r.URL.Path)
		h.ServeHTTP(w, r)
	})
}
