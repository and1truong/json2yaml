package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func main() {
	rootCmd := newRootCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "json2yaml [file]",
		Short: "Convert JSON to YAML",
		Long: `Convert JSON files or stdin input to YAML format.
Examples:
  json2yaml input.json
  cat input.json | json2yaml`,
		Args: cobra.MaximumNArgs(1),
		RunE: runConvert,
	}
	return cmd
}

func runConvert(cmd *cobra.Command, args []string) error {
	var input []byte
	var err error

	if len(args) > 0 {
		input, err = os.ReadFile(args[0])
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}
	} else {
		input, err = io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return fmt.Errorf("error reading stdin: %w", err)
		}
	}

	yamlData, err := convertJSONToYAML(input)
	if err != nil {
		return fmt.Errorf("error converting JSON to YAML: %w", err)
	}

	fmt.Fprint(cmd.OutOrStdout(), string(yamlData))
	return nil
}

func convertJSONToYAML(jsonData []byte) ([]byte, error) {
	var data interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}
	return yaml.Marshal(data)
}
