// cmd/tscli/get/dns/split/cli.go
//
// `tscli get dns split`
// Return the **split-DNS domains** (aka “override domains”) configured
// for the current tailnet.
package split

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
)

// Command registers `tscli get splitdns`
func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "split",
		Short: "Get tailnet split DNS domains",
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
				"/tailnet/{tailnet}/dns/split-dns",
				nil,
				&raw,
			); err != nil {
				return fmt.Errorf("failed to fetch split-DNS list: %w", err)
			}

			pretty, _ := json.MarshalIndent(raw, "", "  ")
			fmt.Fprintln(os.Stdout, string(pretty))
			return nil
		},
	}
}
