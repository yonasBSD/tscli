// cmd/tscli/devices/expire/command.go

package expire

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "expire",
		Short: "Expire (invalidate) a device key",
		Long:  "Call the Tailscale API to mark a device's key as expired.",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			deviceID, err := cmd.Flags().GetString("device")
			if err != nil {
				return fmt.Errorf("failed to read --device flag: %w", err)
			}

			if _, err = tscli.Do(
				cmd.Context(),
				client,
				http.MethodPost,
				"/device/"+deviceID+"/expire",
				nil, // no request body
				nil, // no response body expected
			); err != nil {
				return err
			}

			// Print a simple JSON confirmation to stdout.
			payload := map[string]string{"result": fmt.Sprintf("device %s expired", deviceID)}
			out, _ := json.MarshalIndent(payload, "", "  ")
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}

	cmd.Flags().String("device", "", "Device ID to expire")
	_ = cmd.MarkFlagRequired("device")
	return cmd
}
