package middleware

import (
	"context"
	"net/http"

	logHelper "stoneBanking/app/domain/entities/logger"
	"stoneBanking/app/domain/entities/token"
	"stoneBanking/app/domain/types"

	"github.com/google/uuid"
)

type Middleware struct {
	l logHelper.Logger
	t token.Authenticator
}

const RequestContextID = types.ContextKey("request_id")

func NewMiddleware(log logHelper.Logger, token token.Authenticator) *Middleware {
	return &Middleware{
		l: log,
		t: token,
	}
}

func (m *Middleware) LogRoutes(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		newCtx := context.WithValue(ctx, RequestContextID, uuid.New().String())
		m.l.SetRequestIDFromContext(newCtx)
		m.l.LogInfo("Middleware.LogRoutes", "received request in url: "+r.URL.Path)
		h.ServeHTTP(w, r.WithContext(newCtx))
	})
}

func GetRequestIDFromContext(ctx context.Context) (string, error) {
	requestID := ctx.Value(RequestContextID)
	value, _ := requestID.(string)
	return value, nil
}
