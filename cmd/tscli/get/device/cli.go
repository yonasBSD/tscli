// cmd/tscli/get/device/cli.go
//
// Fetch details for a single device.
//
// Examples
//
//	# Standard fields
//	tscli get device --device node-abcdef123456
//
//	# Every field the API can return
//	tscli get device --device node-abcdef123456 --all
package device

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	tsapi "tailscale.com/client/tailscale/v2"
)

func Command() *cobra.Command {
	var (
		showAll  bool
		deviceID string
	)

	cmd := &cobra.Command{
		Use:   "device",
		Short: "Get a device's information",
		Long:  "Return a single device record from the Tailscale API.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if deviceID == "" {
				return fmt.Errorf("--device is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			var d *tsapi.Device
			if showAll {
				d, err = client.Devices().GetWithAllFields(cmd.Context(), deviceID)
			} else {
				d, err = client.Devices().Get(cmd.Context(), deviceID)
			}
			if err != nil {
				return fmt.Errorf("failed to get device: %w", err)
			}

			out, err := json.MarshalIndent(d, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal JSON: %w", err)
			}
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}

	cmd.Flags().BoolVar(
		&showAll,
		"all",
		false,
		"Include advanced fields such as ClientConnectivity, AdvertisedRoutes, and EnabledRoutes (equivalent to '?fields=all').",
	)
	cmd.Flags().StringVar(
		&deviceID,
		"device",
		"",
		`Device identifier to query (nodeId "node-abcdef123456" or numeric id).`,
	)
	_ = cmd.MarkFlagRequired("device")

	return cmd
}
