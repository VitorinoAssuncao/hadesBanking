package token

import (
	"stoneBanking/app/domain/entities/token"
)

type tokenRepository struct {
	signKey string
}

func NewTokenRepository(signKey string) token.TokenInterface {
	return &tokenRepository{
		signKey: signKey,
	}
}
