// cmd/tscli/set/policy/cli.go
package policyset

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/tailscale/hujson"
)

func Command() *cobra.Command {
	var (
		filePath string
		inline   string
		ifMatch  string
	)

	cmd := &cobra.Command{
		Use:   "policy",
		Short: "Replace the tailnet ACL / policy file",
		Long: `Accepts JSON or HuJSON.

  --file  file://path.hujson  (or "-": stdin)
  --body  '{"ACLs":[...]}'   (inline)

  Optional optimistic lock:
  --if-match ts-default  OR  value from previous GET /acl ETag`,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if filePath == "" && inline == "" {
				return errors.New("one of --file or --body is required")
			}
			if filePath != "" && inline != "" {
				return errors.New("--file and --body are mutually exclusive")
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, _ []string) error {
			raw, err := readInput(filePath, inline)
			if err != nil {
				return err
			}

			if err := validateJSONorHUJSON(raw); err != nil {
				return fmt.Errorf("invalid ACL: %w", err)
			}

			client, err := tscli.New()
			if err != nil {
				return err
			}

			if err := client.PolicyFile().Set(cmd.Context(), string(raw), ifMatch); err != nil {
				return fmt.Errorf("update failed: %w", err)
			}

			updated, err := client.PolicyFile().Raw(cmd.Context())
			if err != nil {
				return fmt.Errorf("fetch updated file: %w", err)
			}
			fmt.Fprintln(os.Stdout, updated.HuJSON) // unchanged formatting
			return nil
		},
	}

	cmd.Flags().StringVar(&filePath, "file", "", "file://path, local path, or '-' for stdin")
	cmd.Flags().StringVar(&inline,   "body", "", "Raw ACL JSON/HuJSON string")
	cmd.Flags().StringVar(&ifMatch,  "if-match", "", "Value for If-Match header (etag or ts-default)")

	return cmd
}

func readInput(path, inline string) ([]byte, error) {
	if inline != "" {
		return []byte(inline), nil
	}

	if path == "-" {
		return io.ReadAll(os.Stdin)
	}
	if strings.HasPrefix(path, "file://") {
		path = strings.TrimPrefix(path, "file://")
	}
	return os.ReadFile(path)
}

func validateJSONorHUJSON(b []byte) error {
	if json.Valid(b) {
		return nil
	}
	_, err := hujson.Parse(b)
	return err
}
