package middleware

import (
	"net/http"
	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	customError "stoneBanking/app/domain/errors"
)

type Middleware struct {
	l logHelper.Logger
}

func NewMiddleware(log logHelper.Logger) *Middleware {
	return &Middleware{
		l: log,
	}
}

func GetAccountIDFromToken(r *http.Request, t token.Authenticator) (string, error) {
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

func (m *Middleware) LogRoutes(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.l.LogInfo("request", "received request in url: "+r.URL.Path)
		h.ServeHTTP(w, r)
	})
}
