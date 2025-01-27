package cmd

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstallationsCmd(t *testing.T) {
	tests := []struct {
		name              string
		args              []string
		mockClient        *mockGithubClient
		mockJWTToken      string
		mockJWTTokenError error
		expectedOutput    string
		expectedError     bool
	}{
		{
			name: "successful list installations",
			args: []string{"--app-id", "123", "--private-key", "test-key"},
			mockClient: &mockGithubClient{
				mockListInstallations: func() (string, error) {
					return `{"installations":[{"id":1}]}`, nil
				},
			},
			mockJWTToken:      "mock-jwt-token",
			mockJWTTokenError: nil,
			expectedOutput:    `{"installations":[{"id":1}]}` + "\n",
			expectedError:     false,
		},
		{
			name: "list installations fails",
			args: []string{"--app-id", "123", "--private-key", "test-key"},
			mockClient: &mockGithubClient{
				mockListInstallations: func() (string, error) {
					return "", errors.New("api error")
				},
			},
			mockJWTToken:      "mock-jwt-token",
			mockJWTTokenError: nil,
			expectedOutput:    "",
			expectedError:     true,
		},
		{
			name:              "jwt token creation fails",
			args:              []string{"--app-id", "123", "--private-key", "test-key"},
			mockClient:        &mockGithubClient{},
			mockJWTToken:      "",
			mockJWTTokenError: errors.New("invalid key"),
			expectedOutput:    "",
			expectedError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := new(bytes.Buffer)
			stderr := new(bytes.Buffer)

			cmd := newInstallationsCmd(
				mockGithubClientFactory(tt.mockClient),
				mockJWTTokenFactory(tt.mockJWTToken, tt.mockJWTTokenError),
			)
			cmd.SetArgs(tt.args)
			cmd.SetOut(stdout)
			cmd.SetErr(stderr)

			err := cmd.Execute()
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, stdout.String())
			}
		})
	}
}
