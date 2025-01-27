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
		{"NoArgs", []string{}, false},
		{"Help", []string{"--help"}, false},
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
