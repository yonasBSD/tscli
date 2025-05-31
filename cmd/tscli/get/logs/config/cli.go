// cmd/tscli/get/logs/config/cli.go
//
// Fetch configuration audit logs from the Tailscale API.
//
//	# Get configuration logs for a specific day
//	tscli get logs config --start 2024-01-01T00:00:00Z --end 2024-01-01T23:59:59Z
//
//	# Get configuration logs for the last hour
//	tscli get logs config --start $(date -d '1 hour ago' -Iseconds) --end $(date -Iseconds)
package config

import (
	"context"
	"encoding/json"

	"net/http"
	"net/url"
	"github.com/jaxxstorm/tscli/pkg/output"
	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Command() *cobra.Command {
	var (
		startTime string
		endTime   string
	)

	cmd := &cobra.Command{
		Use:   "config",
		Short: "Get configuration audit logs for the tailnet",
		Long:  "List all configuration audit logs for a tailnet within a time range.",

		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

			path := "/tailnet/{tailnet}/logging/configuration"
			q := url.Values{}
			q.Add("start", startTime)
			q.Add("end", endTime)
			path += "?" + q.Encode()

			var resp json.RawMessage
			if _, err := tscli.Do(
				context.Background(),
				client,
				http.MethodGet,
				path,
				nil,
				&resp,
			); err != nil {
				return err
			}

			out, _ := json.MarshalIndent(resp, "", "  ")
			format := viper.GetString("format")
			output.Print(format, out)
			return nil
		},
	}

	cmd.Flags().StringVarP(&startTime, "start", "s", "",
		`Start time in RFC3339 format (required). Example: "2024-01-01T00:00:00Z"`)
	cmd.Flags().StringVarP(&endTime, "end", "e", "",
		`End time in RFC3339 format (required). Example: "2024-01-01T23:59:59Z"`)

	_ = cmd.MarkFlagRequired("start")
	_ = cmd.MarkFlagRequired("end")

	return cmd
}
