package middleware

import (
	"context"
	"net/http"
	customError "stoneBanking/app/domain/errors"
	"stoneBanking/app/gateway/http/response"
	"stoneBanking/app/gateway/http/transfer/vo/output"
)

func (m *Middleware) GetAccountIDFromTokenLogRoutes(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := response.NewResponse(w)
		headerToken := r.Header.Get("Authorization")
		if headerToken == "" {
			resp.Unauthorized(output.NewError(customError.ErrorServerTokenNotFound))
			return
		}
		accountExternalID, err := m.t.ExtractAccountIDFromToken(headerToken)
		if err != nil {
			resp.BadRequest(output.NewError(err))
			return
		}

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
