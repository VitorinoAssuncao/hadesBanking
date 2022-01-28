package middleware

import (
	"context"
	"net/http"
)

func (m *Middleware) GetAccountIDFromTokenLogRoutes(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")
		/*		if headerToken == "" {
					return "", customError.ErrorServerTokenNotFound
				}
		*/
		accountExternalID, _ := m.t.ExtractAccountIDFromToken(headerToken)

		ctx := r.Context()
		newCtx := context.WithValue(ctx, AccountContextKey, accountExternalID)
		h.ServeHTTP(w, r.WithContext(newCtx))
	})
}

func GetAccountIDFromContext(ctx context.Context) (string, error) {
	accountID := ctx.Value(AccountContextKey)
	value, _ := accountID.(string)
	return value, nil
}
