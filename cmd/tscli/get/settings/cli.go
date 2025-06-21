// cmd/tscli/get/settings/cli.go
//
// `tscli get settings`
// Fetch the tailnet-wide settings object and print it as JSON.
package settings

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
	return &cobra.Command{
		Use:   "settings",
		Short: "Get tailnet-wide settings",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			var s *tsapi.TailnetSettings
			s, err = client.TailnetSettings().Get(context.Background())
			if err != nil {
				return fmt.Errorf("failed to retrieve settings: %w", err)
			}

			out, err := json.MarshalIndent(s, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal settings: %w", err)
			}
			outputType := viper.GetString("output")
			output.Print(outputType, out)
			return nil
		},
	}
}
