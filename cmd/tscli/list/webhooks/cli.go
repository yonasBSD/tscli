// cmd/tscli/list/webhooks/cli.go
package webhooks

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
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}
}
