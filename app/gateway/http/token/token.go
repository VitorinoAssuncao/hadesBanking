package token

import (
	"stoneBanking/app/domain/entities/token"
)

type TokenAuthenticator struct {
	signKey string
}

func NewTokenAuthenticator(signKey string) token.Authenticator {
	return &TokenAuthenticator{
		signKey: signKey,
	}
}
