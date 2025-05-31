// cmd/tscli/get/dns/nameservers/cli.go
//
// `tscli get ns`
// Return the list of custom DNS nameservers configured for the current tailnet.
package nameservers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jaxxstorm/tscli/pkg/output"
	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

			out, _ := json.MarshalIndent(raw, "", "  ")
			format := viper.GetString("format")
			output.Print(format, out)
			return nil
		},
	}
}
