package get

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	tsapi "tailscale.com/client/tailscale/v2"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "get",
		Short: "Get device commands",
		Long:  "Get commands that operate on devices",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			all, err := cmd.Flags().GetBool("all")

			if err != nil {
				return fmt.Errorf("failed to get all flag: %w", err)
			}

			deviceID, err := cmd.Flags().GetString("device")
			if err != nil {
				return fmt.Errorf("failed to get device flag: %w", err)
			}

			var device *tsapi.Device

			if all {
				device, err = client.Devices().GetWithAllFields(context.Background(), deviceID)
				if err != nil {
					return fmt.Errorf("failed to list devices with all fields: %w", err)
				}
			} else {
				device, err = client.Devices().Get(context.Background(), deviceID)
				if err != nil {
					return fmt.Errorf("failed to list devices: %w", err)
				}
			}

			out, err := json.MarshalIndent(device, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal devices into JSON: %w", err)
			}
			fmt.Fprintln(os.Stdout, string(out))
			return nil

		},
	}

	command.Flags().Bool("all", false, "Display all fields in result.")
	command.Flags().String("device", "", "Device ID to get.")
	_ = command.MarkFlagRequired("device")

	return command
}
