package contacts

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
	var contactID string

	cmd := &cobra.Command{
		Use:   "contacts",
		Short: "Get tailnet contacts from the Tailscale API",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

			var c *tsapi.Contacts
			c, err = client.Contacts().Get(context.Background())
			if err != nil {
				return fmt.Errorf("failed to get contact %s: %w", contactID, err)
			}

			out, _ := json.MarshalIndent(c, "", "  ")
			outputType := viper.GetString("output")
			output.Print(outputType, out)
			return nil
		},
	}

	return cmd
}
