// cmd/tscli/get/key/cli.go
//
// `tscli get key --key <auth-key-id>`
// Fetch details for a single reusable auth-key.

package key

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
	var keyID string

	cmd := &cobra.Command{
		Use:   "key",
		Short: "Get a single tailnet auth key",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			var k *tsapi.Key
			k, err = client.Keys().Get(context.Background(), keyID)
			if err != nil {
				return fmt.Errorf("failed to get key %s: %w", keyID, err)
			}

			out, err := json.MarshalIndent(k, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal key into JSON: %w", err)
			}
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}

	cmd.Flags().StringVar(&keyID, "key", "", "Key ID to retrieve")
	_ = cmd.MarkFlagRequired("key")

	return cmd
}
