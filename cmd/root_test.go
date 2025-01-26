package cmd

import (
	"bytes"
	"testing"
)

func TestRootCmd(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"NoArgs", []string{}, true},
		{"Help", []string{"--help"}, false},
		{"Create", []string{"create", "--app-id=123", "--private-key=key"}, false},
		{"Revoke", []string{"revoke", "--app-id=123", "--private-key=key"}, false},
		{"Installations", []string{"installations", "--app-id=123", "--private-key=key"}, false},
		{"MissingPrivateKey", []string{"create", "--app-id=123"}, true},
		{"MissingAppID", []string{"create", "--private-key=key"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdtout := new(bytes.Buffer)
			stderr := new(bytes.Buffer)
			err := Execute(tt.args, sdtout, stderr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
