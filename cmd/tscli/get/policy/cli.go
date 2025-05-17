// cmd/tscli/get/policy/cli.go
package policy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jaxxstorm/tscli/cmd/tscli/get/policy/preview"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/tailscale/hujson"
)

func Command() *cobra.Command {
	var asJSON bool

	cmd := &cobra.Command{
		Use:   "policy",
		Short: "Print the tailnet policy file (HUJSON or JSON)",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			policy, err := client.PolicyFile().Raw(cmd.Context())
			if err != nil {
				return fmt.Errorf("failed to retrieve policy file: %w", err)
			}

			if _, err := hujson.Parse([]byte(policy.HuJSON)); err != nil {
				return fmt.Errorf("failed to parse HuJSON policy file: %w", err)
			}

			if asJSON {
				std, err := hujson.Standardize([]byte(policy.HuJSON))
				if err != nil {
					return fmt.Errorf("failed to convert to JSON: %w", err)
				}

				var pretty bytes.Buffer
				if err := json.Indent(&pretty, std, "", "  "); err != nil {
					return err
				}
				fmt.Fprintln(os.Stdout, pretty.String())
			} else {
				fmt.Fprintln(os.Stdout, policy.HuJSON)
			}
			return nil
		},
	}

	cmd.AddCommand(preview.Command())

	cmd.Flags().BoolVar(&asJSON, "json", false, "Output the policy as canonical JSON instead of HUJSON")
	return cmd
}
