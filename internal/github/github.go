package github

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/google/go-github/v68/github"
)

type githubClientImpl struct {
	client *github.Client
}

func NewClient(token string) GithubClient {
	return &githubClientImpl{
		client: github.NewClient(nil).WithAuthToken(token),
	}
}

func (g *githubClientImpl) GetInstallationID() (int64, error) {
	installations, _, err := g.client.Apps.ListInstallations(context.Background(), &github.ListOptions{})
	if err != nil {
		return 0, err
	}
	if len(installations) == 0 {
		return 0, errors.New("no installations found")
	}
	return installations[0].GetID(), nil
}

func (g *githubClientImpl) CreateInstallationToken(installationID int64) (string, error) {
	installationToken, _, err := g.client.Apps.CreateInstallationToken(context.Background(), installationID, &github.InstallationTokenOptions{})
	if err != nil {
		return "", err
	}
	return installationToken.GetToken(), nil
}

func (g *githubClientImpl) RevokeInstallationToken() error {
	_, err := g.client.Apps.RevokeInstallationToken(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (g *githubClientImpl) ListInstallations() (string, error) {
	installations, _, err := g.client.Apps.ListInstallations(context.Background(), &github.ListOptions{})
	if err != nil {
		return "", err
	}
	jsonBytes, err := json.MarshalIndent(installations, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
