// cmd/tscli/get/policy/cli.go
package policy

import (
	"fmt"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/tailscale/hujson"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "policy",
		Short: "Pretty-print the tailnet policy file (HUJSON)",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			policy, err := client.PolicyFile().Raw(cmd.Context())
			if err != nil {
				return fmt.Errorf("failed to retrieve policy file: %w", err)
			}

			parsed, err := hujson.Parse([]byte(policy.HuJSON))
			if err != nil {
				return fmt.Errorf("failed to parse HuJSON policy file: %w", err)
			}



			fmt.Fprintln(os.Stdout, parsed)
			return nil
		},
	}
}
