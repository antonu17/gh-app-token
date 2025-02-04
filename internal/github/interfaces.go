package github

type GithubClient interface {
	GetInstallationID() (int64, error)
	CreateInstallationToken(int64) (string, error)
	RevokeInstallationToken() error
	ListInstallations() (string, error)
}

var _ GithubClient = &githubClientImpl{}
