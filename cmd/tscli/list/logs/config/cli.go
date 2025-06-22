// cmd/tscli/list/logs/config/cli.go
//
// Fetch configuration-audit logs.
//
//	# Last 12 h
//	tscli list logs config --start 12h
//
//	# Specific window
//	tscli list logs config -s 2025-05-01T00:00:00Z -e 2025-05-01T23:59:59Z
package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	tstime "github.com/jaxxstorm/tscli/pkg/time"
	"time"

	"github.com/jaxxstorm/tscli/pkg/output"
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
		Use:   "config",
		Short: "Get configuration-audit logs for the tailnet",
		Long:  "Fetch audit-log entries for a given period. --start accepts RFC3339 *or* a relative offset such as 30d, 6h, 90m.",
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

			path := "/tailnet/{tailnet}/logging/configuration"
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

			return output.Print(viper.GetString("output"), raw)
		},
	}

	cmd.Flags().StringVarP(&startFlag, "start", "s", "24h",
		`Start time (RFC3339 or relative like "24h", "30d12h"). Required.`)
	cmd.Flags().StringVarP(&endFlag, "end", "e", "",
		`End time in RFC3339. Defaults to “now”.`)

	return cmd
}
