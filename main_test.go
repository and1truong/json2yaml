package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestCLI(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		input       string
		wantOutput  string
		wantErr     bool
		wantErrMsg  string
	}{
		{
			name:       "valid file input",
			args:       []string{"testdata/valid.json"},
			wantOutput: "key: value\n",
		},
		{
			name:       "valid stdin input",
			input:      `{"key": "value"}`,
			wantOutput: "key: value\n",
		},
		{
			name:        "invalid json file",
			args:        []string{"testdata/invalid.json"},
			wantErr:     true,
			wantErrMsg:  "error converting JSON to YAML",
		},
		{
			name:        "file not found",
			args:        []string{"nonexistent.json"},
			wantErr:     true,
			wantErrMsg:  "error reading file",
		},
		{
			name:        "invalid stdin json",
			input:       `{invalid}`,
			wantErr:     true,
			wantErrMsg:  "error converting JSON to YAML",
		},
		{
			name:        "empty file",
			args:        []string{"testdata/empty.json"},
			wantErr:     true,
			wantErrMsg:  "error converting JSON to YAML",
		},
		{
			name:        "empty stdin",
			input:       "",
			wantErr:     true,
			wantErrMsg:  "error converting JSON to YAML",
		},
	}

	// Create testdata directory with sample files
	os.Mkdir("testdata", 0755)
	defer os.RemoveAll("testdata")
	createTestFile(t, "testdata/valid.json", `{"key": "value"}`)
	createTestFile(t, "testdata/invalid.json", `{invalid}`)
	createTestFile(t, "testdata/empty.json", "")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create root command
			rootCmd := newRootCommand()
			
			// Capture output
			var stdout, stderr bytes.Buffer
			rootCmd.SetOut(&stdout)
			rootCmd.SetErr(&stderr)
			
			// Set up input
			if tt.input != "" {
				rootCmd.SetIn(strings.NewReader(tt.input))
			}

			// Execute command
			rootCmd.SetArgs(tt.args)
			err := rootCmd.Execute()

			// Verify results
			if (err != nil) != tt.wantErr {
				t.Errorf("Unexpected error status: %v", err)
			}

			if tt.wantErr {
				if !strings.Contains(stderr.String(), tt.wantErrMsg) {
					t.Errorf("Expected error containing %q, got %q", 
						tt.wantErrMsg, stderr.String())
				}
			} else {
				if strings.TrimSpace(stdout.String()) != strings.TrimSpace(tt.wantOutput) {
					t.Errorf("Output mismatch\nGot: %q\nWant: %q", 
						stdout.String(), tt.wantOutput)
				}
			}
		})
	}
}

func createTestFile(t *testing.T, path, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
}

// Helper to create command instance
func newRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "json2yaml [file]",
		Short: "Convert JSON to YAML",
		RunE:  runConvert,
	}
	rootCmd.SetOut(io.Discard)
	rootCmd.SetErr(io.Discard)
	return rootCmd
}
