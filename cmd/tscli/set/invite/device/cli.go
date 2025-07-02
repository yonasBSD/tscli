// cmd/tscli/set/invite/device/cli.go
//
// `tscli set invite device --id <invite-id> --status <resend|accept>`
// Resend or accept a device invite via POST /api/v2/tailnet/{tailnet}/device-invites/{id}/{action}

package device

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jaxxstorm/tscli/pkg/output"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var validStatuses = map[string]string{
	"resend": "resend",
	"accept": "accept",
}

func Command() *cobra.Command {
	var (
		inviteID string
		status   string
	)

	cmd := &cobra.Command{
		Use:   "device",
		Short: "Set device invite status",
		Long:  "Resend or accept a device invite. Valid statuses: resend, accept",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if inviteID == "" {
				return fmt.Errorf("--id is required")
			}
			if status == "" {
				return fmt.Errorf("--status is required")
			}
			if _, ok := validStatuses[strings.ToLower(status)]; !ok {
				return fmt.Errorf("invalid --status: %s (resend|accept)", status)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			action := validStatuses[strings.ToLower(status)]
			endpoint := fmt.Sprintf("/tailnet/{tailnet}/device-invites/%s/%s", inviteID, action)

			var response map[string]interface{}
			if _, err := tscli.Do(
				context.Background(),
				client,
				http.MethodPost,
				endpoint,
				nil, // no request body
				&response,
			); err != nil {
				return fmt.Errorf("failed to %s device invite %s: %w", action, inviteID, err)
			}

			// If the API returns data, show it, otherwise show confirmation
			if len(response) > 0 {
				out, _ := json.MarshalIndent(response, "", "  ")
				outputType := viper.GetString("output")
				output.Print(outputType, out)
			} else {
				// Fallback confirmation message
				payload := map[string]string{
					"result": fmt.Sprintf("device invite %s %sed", inviteID, action),
					"action": action,
					"id":     inviteID,
				}
				out, _ := json.MarshalIndent(payload, "", "  ")
				outputType := viper.GetString("output")
				output.Print(outputType, out)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&inviteID, "id", "", "Device invite ID")
	cmd.Flags().StringVar(&status, "status", "", "Action to perform (resend|accept)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("status")

	return cmd
}
