package middleware

import (
	"net/http"

	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	"stoneBanking/app/domain/types"
)

type Middleware struct {
	l logHelper.Logger
	t token.Authenticator
}

const AccountContextKey = types.ContextKey("account_id")

func NewMiddleware(log logHelper.Logger, token token.Authenticator) *Middleware {
	return &Middleware{
		l: log,
		t: token,
	}
}

func (m *Middleware) LogRoutes(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.l.LogInfo("request", "received request in url: "+r.URL.Path)
		h.ServeHTTP(w, r)
	})
}
