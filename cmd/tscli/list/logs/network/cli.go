// cmd/tscli/list/logs/network/cli.go
//
// Fetch network-audit logs.
//
//	# The last 30 minutes
//	tscli list logs network --start 30m
//
//	# A specific day
//	tscli list logs network -s 2024-05-01T00:00:00Z -e 2024-05-01T23:59:59Z
package network

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/jaxxstorm/tscli/pkg/output"
	tstime "github.com/jaxxstorm/tscli/pkg/time"
	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


func Command() *cobra.Command {
	var (
		startFlag string
		endFlag   string
	)

	cmd := &cobra.Command{
		Use:   "network",
		Short: "Get network audit logs for the tailnet",
		Long:  "Fetch audit log entries for a given period. --start accepts RFC3339 or a relative offset like 30d, 12h, 45m.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			start, err := tstime.ParseTime(startFlag, true)
			if err != nil {
				return fmt.Errorf("invalid --start: %w", err)
			}

			end, err := tstime.ParseTime(endFlag, true)
			if err != nil {
				return fmt.Errorf("invalid --end: %w", err)
			}
			if !start.Before(end) {
				return errors.New("--start must be before --end/now")
			}

			client, err := tscli.New()
			if err != nil {
				return err
			}

			path := "/tailnet/{tailnet}/logging/network"
			q := url.Values{}
			q.Set("start", start.Format(time.RFC3339))
			q.Set("end", end.Format(time.RFC3339))
			path += "?" + q.Encode()

			var raw json.RawMessage
			if _, err := tscli.Do(
				context.Background(), client, http.MethodGet,
				path, nil, &raw,
			); err != nil {
				return err
			}

			return output.Print(viper.GetString("format"), raw)
		},
	}

	cmd.Flags().StringVarP(&startFlag, "start", "s", "90m",
		`RFC3339 timestamp *or* relative offset (e.g. "30d", "90m"). Required.`)
	cmd.Flags().StringVarP(&endFlag, "end", "e", "",
		`RFC3339 timestamp. Defaults to the current time.`)

	return cmd
}
