package cmd

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCmd(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		mockClient   *mockGithubClient
		mockJWTToken string
		mockJWTErr   error
		wantOutput   string
		wantErr      bool
		wantErrMsg   string
	}{
		{
			name: "success - creates installation token",
			args: []string{"--app-id", "123", "--private-key", "test-key"},
			mockClient: &mockGithubClient{
				mockGetInstallationID: func() (int64, error) {
					return 456, nil
				},
				mockCreateInstallationToken: func(id int64) (string, error) {
					return "test-installation-token", nil
				},
			},
			mockJWTToken: "test-jwt-token",
			mockJWTErr:   nil,
			wantOutput:   "test-installation-token\n",
			wantErr:      false,
		},
		{
			name:         "error - jwt token creation fails",
			args:         []string{"--app-id", "123", "--private-key", "test-key"},
			mockClient:   &mockGithubClient{},
			mockJWTToken: "",
			mockJWTErr:   errors.New("jwt error"),
			wantOutput:   "",
			wantErr:      true,
			wantErrMsg:   "error generating JWT token for Github App: jwt error",
		},
		{
			name: "error - get installation ID fails",
			args: []string{"--app-id", "123", "--private-key", "test-key"},
			mockClient: &mockGithubClient{
				mockGetInstallationID: func() (int64, error) {
					return 0, errors.New("installation error")
				},
			},
			mockJWTToken: "test-jwt-token",
			mockJWTErr:   nil,
			wantOutput:   "",
			wantErr:      true,
			wantErrMsg:   "error getting installation id: installation error",
		},
		{
			name: "error - create installation token fails",
			args: []string{"--app-id", "123", "--private-key", "test-key"},
			mockClient: &mockGithubClient{
				mockGetInstallationID: func() (int64, error) {
					return 456, nil
				},
				mockCreateInstallationToken: func(id int64) (string, error) {
					return "", errors.New("token error")
				},
			},
			mockJWTToken: "test-jwt-token",
			mockJWTErr:   nil,
			wantOutput:   "",
			wantErr:      true,
			wantErrMsg:   "error generating installation token for Github App: token error",
		},
		{
			name:       "error - missing app-id flag",
			args:       []string{"--private-key", "test-key"},
			mockClient: &mockGithubClient{},
			wantOutput: "",
			wantErr:    true,
		},
		{
			name:       "error - missing private-key flag",
			args:       []string{"--app-id", "123"},
			mockClient: &mockGithubClient{},
			wantOutput: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := new(bytes.Buffer)
			stderr := new(bytes.Buffer)

			cmd := newCreateCmd(
				mockGithubClientFactory(tt.mockClient),
				mockJWTTokenFactory(tt.mockJWTToken, tt.mockJWTErr),
			)
			cmd.SetArgs(tt.args)
			cmd.SetOut(stdout)
			cmd.SetErr(stderr)

			err := cmd.Execute()
			if tt.wantErr {
				assert.Error(t, err)
				if tt.wantErrMsg != "" {
					assert.Equal(t, tt.wantErrMsg, err.Error())
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantOutput, stdout.String())
			}
		})
	}
}
