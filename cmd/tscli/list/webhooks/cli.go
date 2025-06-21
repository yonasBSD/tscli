// cmd/tscli/list/webhooks/cli.go
package webhooks

import (
	"context"
	"encoding/json"

	"github.com/jaxxstorm/tscli/pkg/output"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tsapi "tailscale.com/client/tailscale/v2"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "webhooks",
		Short: "List tailnet webhooks",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

			var webhooks []tsapi.Webhook
			webhooks, err = client.Webhooks().List(context.Background())
			if err != nil {
				return err
			}

			out, _ := json.MarshalIndent(webhooks, "", "  ")
			outputType := viper.GetString("output")
			output.Print(outputType, out)
			return nil
		},
	}
}
