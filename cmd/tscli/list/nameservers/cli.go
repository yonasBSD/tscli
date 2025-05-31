// cmd/tscli/list/nameservers/cli.go
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

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "nameservers",
		Short: "List custom DNS nameservers for the tailnet",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

			var raw json.RawMessage // <- receives the body untouched
			if _, err := tscli.Do(
				context.Background(),
				client,
				http.MethodGet,
				"/tailnet/{tailnet}/dns/nameservers",
				nil,
				&raw,
			); err != nil {
				return fmt.Errorf("failed to list DNS nameservers: %w", err)
			}

			out, _ := json.MarshalIndent(raw, "", "  ")
			format := viper.GetString("format")
			output.Print(format, out)
			return nil
		},
	}
}
