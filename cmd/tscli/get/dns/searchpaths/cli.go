// cmd/tscli/get/dns/searchpaths/cli.go
//
// `tscli get dns searchpaths`
// Return the list of DNS *search domains* (a k a search-paths) configured
// for the current tailnet.
package searchpaths

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
)

// Command registers `tscli get searchpaths`
func Command() *cobra.Command {
	return &cobra.Command{
		Use:     "searchpaths",
		Aliases: []string{"sp"},
		Short:   "Get tailnet DNS search-paths",
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
				"/tailnet/{tailnet}/dns/searchpaths",
				nil,
				&raw,
			); err != nil {
				return fmt.Errorf("failed to fetch DNS search-paths: %w", err)
			}

			pretty, _ := json.MarshalIndent(raw, "", "  ")
			fmt.Fprintln(os.Stdout, string(pretty))
			return nil
		},
	}
}
