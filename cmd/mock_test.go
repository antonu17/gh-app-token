package cmd

import (
	"github.com/n26/gh-app-token/internal/github"
)

type mockGithubClient struct {
	mockGetInstallationID       func() (int64, error)
	mockCreateInstallationToken func(int64) (string, error)
	mockRevokeInstallationToken func() error
}

func (m *mockGithubClient) GetInstallationID() (int64, error) {
	return m.mockGetInstallationID()
}

func (m *mockGithubClient) CreateInstallationToken(id int64) (string, error) {
	return m.mockCreateInstallationToken(id)
}

func (m *mockGithubClient) RevokeInstallationToken() error {
	return m.mockRevokeInstallationToken()
}

func mockGithubClientFactory(client github.GithubClient) GithubClientFactory {
	return func(string) github.GithubClient {
		return client
	}
}

func mockJWTTokenFactory(token string, err error) JWTTokenFactory {
	return func(issuer string, privateKey string) (string, error) {
		return token, err
	}
}

type cmdTestCase struct {
	name           string
	args           []string
	mockClient     *mockGithubClient
	expectedOutput string
	expectedError  bool
}
