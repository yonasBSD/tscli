// cmd/tscli/get/key/cli.go
//
// `tscli get key --key <auth-key-id>`
// Fetch details for a single reusable auth-key.

package key

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jaxxstorm/tscli/pkg/output"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

			var raw *tsapi.Key
			raw, err = client.Keys().Get(context.Background(), keyID)
			if err != nil {
				return fmt.Errorf("failed to get key %s: %w", keyID, err)
			}

			out, _ := json.MarshalIndent(raw, "", "  ")
			outputType := viper.GetString("output")
			output.Print(outputType, out)
			return nil
		},
	}

	cmd.Flags().StringVar(&keyID, "key", "", "Key ID to retrieve")
	_ = cmd.MarkFlagRequired("key")

	return cmd
}
