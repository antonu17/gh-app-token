package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCmd(t *testing.T) {
	tests := []cmdTestCase{
		{
			name: "successful token creation",
			args: []string{"create", "--app-id", "123", "--private-key", "test-key"},
			mockClient: &mockGithubClient{
				mockGetInstallationID: func() (int64, error) {
					return 456, nil
				},
				mockCreateInstallationToken: func(id int64) (string, error) {
					return "test-token", nil
				},
			},
			expectedOutput: "test-token\n",
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := new(bytes.Buffer)
			stderr := new(bytes.Buffer)

			rootCmd := newCreateCmd(mockGithubClientFactory(tt.mockClient), mockJWTTokenFactory("", nil))
			rootCmd.SetArgs(tt.args)
			rootCmd.SetOut(stdout)
			rootCmd.SetErr(stderr)

			err := rootCmd.Execute()
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, stdout.String())
			}
		})
	}
}
