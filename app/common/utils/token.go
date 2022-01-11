package utils

import (
	customError "stoneBanking/app/domain/errors"

	"github.com/golang-jwt/jwt"
)

type ClaimStruct struct {
	jwt.StandardClaims
	UserID string
}

func GenerateToken(accountExternalID, signingKey string) (signedToken string, err error) {
	mySigningKey := []byte(signingKey)
	claims := &ClaimStruct{
		UserID:         accountExternalID,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ExtractClaims(tokenStr, signingKey string) (string, error) {
	mySigningKey := []byte(signingKey)
	token, err := jwt.ParseWithClaims(tokenStr, &ClaimStruct{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return "", customError.ErrorServerExtractToken
	}

	claims := token.Claims.(*ClaimStruct)
	return claims.UserID, nil
}
