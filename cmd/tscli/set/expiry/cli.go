package expiry

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
	cmd := &cobra.Command{
		Use:   "expiry",
		Short: "Set device key expiry",
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
				context.Background(),
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
			format := viper.GetString("format")
			output.Print(format, out)
			return nil
		},
	}

	cmd.Flags().String("device", "", `Device ID whose key will be expired(nodeId "node-abc123" or numeric id). Example: --device=node-abcdef123456`)
	_ = cmd.MarkFlagRequired("device")
	return cmd
}
