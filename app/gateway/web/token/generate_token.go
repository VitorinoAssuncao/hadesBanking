package token

import "github.com/golang-jwt/jwt"

func (r TokenRepository) GenerateToken(accountExternalID string) (signedToken string, err error) {
	type ClaimStruct struct {
		jwt.StandardClaims
		UserID string
	}

	mySigningKey := []byte(r.signKey)
	claims := ClaimStruct{
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
