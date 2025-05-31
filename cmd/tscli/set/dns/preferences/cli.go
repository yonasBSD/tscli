// cmd/tscli/set/dns/preferences/cli.go
//
// Toggle MagicDNS for the current tailnet.
//
//	# enable MagicDNS
//	tscli set dns prefs --magicdns
//
//	# disable MagicDNS
//	tscli set dns prefs --magicdns=false
package preferences

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

// Command registers `tscli set dns-prefs`
func Command() *cobra.Command {
	var magicDNS bool

	cmd := &cobra.Command{
		Use:     "prefs",
		Aliases: []string{"preferences"},
		Short:   "Enable or disable MagicDNS for the tailnet",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

			body := map[string]bool{"magicDNS": magicDNS}

			var raw json.RawMessage
			if _, err := tscli.Do(
				context.Background(),
				client,
				http.MethodPost,
				"/tailnet/{tailnet}/dns/preferences",
				body,
				&raw,
			); err != nil {
				return fmt.Errorf("update failed: %w", err)
			}

			out, _ := json.MarshalIndent(raw, "", "  ")
			format := viper.GetString("format")
			output.Print(format, out)
			return nil

		},
	}

	cmd.Flags().BoolVar(&magicDNS, "magicdns", true,
		"Set to true to enable MagicDNS, false to disable")

	return cmd
}
