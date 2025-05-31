// cmd/tscli/get/logs/network/cli.go
//
// Fetch network audit logs from the Tailscale API.
//
//	# Get network logs for a specific day
//	tscli get logs network --start 2024-01-01T00:00:00Z --end 2024-01-01T23:59:59Z
//
//	# Get network logs for the last hour
//	tscli get logs network --start $(date -d '1 hour ago' -Iseconds) --end $(date -Iseconds)
package network

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/jaxxstorm/tscli/pkg/output"
)

func Command() *cobra.Command {
	var (
		startTime string
		endTime   string
	)

	cmd := &cobra.Command{
		Use:   "network",
		Short: "Get network audit logs for the tailnet",
		Long:  "List all network audit logs for a tailnet within a time range.",

		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

			path := "/tailnet/{tailnet}/logging/network"
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
