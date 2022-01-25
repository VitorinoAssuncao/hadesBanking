package token

import "github.com/golang-jwt/jwt"

func (t TokenAuthenticator) ExtractAccountIDFromToken(tokenStr string) (accountExternalID string, err error) {
	type ClaimStruct struct {
		jwt.StandardClaims
		UserID string
	}

	mySigningKey := []byte(t.signKey)
	token, err := jwt.ParseWithClaims(tokenStr, &ClaimStruct{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return "", err
	}

	claims := token.Claims.(*ClaimStruct)
	return claims.UserID, nil
}
