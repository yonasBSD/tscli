package contacts

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
	var contactID string

	cmd := &cobra.Command{
		Use:   "contacts",
		Short: "Get tailnet contacts from the Tailscale API",
		RunE: func(_ *cobra.Command, _ []string) error {
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
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}

	return cmd
}
