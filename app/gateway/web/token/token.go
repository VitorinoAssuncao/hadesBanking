package token

import (
	"stoneBanking/app/domain/entities/token"
)

type TokenRepository struct {
	signKey string
}

func NewTokenRepository(signKey string) token.Repository {
	return &TokenRepository{
		signKey: signKey,
	}
}
