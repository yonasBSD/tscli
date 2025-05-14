package authorize

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "authorize",
		Short: "Authorize or de-authorize a device",
		Long:  "Approve or reject a Tailscale device by ID.",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create Tailscale client: %w", err)
			}

			approved, err := cmd.Flags().GetBool("approve")
			if err != nil {
				return fmt.Errorf("failed to read --approve flag: %w", err)
			}

			deviceID, err := cmd.Flags().GetString("device")
			if err != nil {
				return fmt.Errorf("failed to read --device flag: %w", err)
			}
			if deviceID == "" {
				return fmt.Errorf("--device flag (device ID) is required")
			}

			if err := client.Devices().SetAuthorized(cmd.Context(), deviceID, approved); err != nil {
				return fmt.Errorf("failed to set authorization: %w", err)
			}

			var payload any
			if approved {
				payload = map[string]string{"result": fmt.Sprintf("device %s is now approved", deviceID)}
			} else {
				payload = map[string]string{"result": fmt.Sprintf("device %s is now unapproved", deviceID)}
			}

			out, err := json.MarshalIndent(payload, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal JSON: %w", err)
			}
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}

	command.Flags().Bool("approve", true, "Approve (true) or reject (false) the device")
	command.Flags().String("device", "", "Device ID to authorize")
	_ = command.MarkFlagRequired("device")

	return command
}
