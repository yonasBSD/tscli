// cmd/tscli/get/nameservers/cli.go
//
// `tscli get ns`
// Return the list of custom DNS nameservers configured for the current tailnet.
package nameservers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
)

// Command registers `tscli get ns`
func Command() *cobra.Command {
	return &cobra.Command{
		Use:     "ns",
		Aliases: []string{"nameservers"},
		Short:   "Get tailnet DNS nameservers",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

			var raw json.RawMessage
			if _, err := tscli.Do(
				context.Background(),
				client,
				http.MethodGet,
				"/tailnet/{tailnet}/dns/nameservers",
				nil,
				&raw,
			); err != nil {
				return fmt.Errorf("failed to fetch nameservers: %w", err)
			}

			pretty, _ := json.MarshalIndent(raw, "", "  ")
			fmt.Fprintln(os.Stdout, string(pretty))
			return nil
		},
	}
}
