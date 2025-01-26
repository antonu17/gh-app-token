package github

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-github/v68/github"
)

func Token() *github.Client {
	return github.NewClient(nil)
}

func NewAppToken(issuer string, privateKey string) (string, error) {
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

func GetInstallationID(appToken string) (int64, error) {
	client := github.NewClient(NewBearerTokenClient(appToken))
	installations, _, err := client.Apps.ListInstallations(context.Background(), &github.ListOptions{})
	if err != nil {
		return 0, err
	}
	if len(installations) == 0 {
		return 0, errors.New("no installations found")
	}
	return installations[0].GetID(), nil
}

func NewInstallationToken(appToken string, installationID int64) (string, error) {
	client := github.NewClient(NewBearerTokenClient(appToken))
	installationToken, _, err := client.Apps.CreateInstallationToken(context.Background(), installationID, &github.InstallationTokenOptions{})
	if err != nil {
		return "", err
	}
	return installationToken.GetToken(), nil
}
