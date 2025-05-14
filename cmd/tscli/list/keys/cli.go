// cmd/tscli/list/keys/cli.go
//
// `tscli list keys`
// List every reusable auth key in the tailnet.

package keys

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	tsapi "tailscale.com/client/tailscale/v2"
)

func Command() *cobra.Command {

	var showAll bool

	cmd := &cobra.Command{
		Use:   "keys",
		Short: "List tailnet auth keys",
		Long:  "Return all reusable authentication keys defined for the tailnet.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			var keys []tsapi.Key
			keys, err = client.Keys().List(context.Background(), showAll)
			if err != nil {
				return fmt.Errorf("failed to list keys: %w", err)
			}

			out, err := json.MarshalIndent(keys, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal keys into JSON: %w", err)
			}
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}

	cmd.Flags().BoolVar(
		&showAll,
		"all",
		false,
		"Include all keys.",
	)

	return cmd
}
