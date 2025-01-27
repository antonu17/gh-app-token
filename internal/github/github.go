package github

import "github.com/google/go-github/v68/github"

type githubClientImpl struct {
	client *github.Client
}

func NewClient(token string) GithubClient {
	return &githubClientImpl{
		client: github.NewClient(nil).WithAuthToken(token),
	}
}
