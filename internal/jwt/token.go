package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(issuer string, privateKey string) (string, error) {
	iat := time.Now().Add(-60 * time.Second)
	exp := time.Now().Add(60 * time.Second)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(iat),
		ExpiresAt: jwt.NewNumericDate(exp),
		Issuer:    issuer,
	})

	rsaKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return "", err
	}

	signedToken, err := jwtToken.SignedString(rsaKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
