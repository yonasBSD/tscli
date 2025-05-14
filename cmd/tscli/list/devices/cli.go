package devices

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	tsapi "tailscale.com/client/tailscale/v2"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "devices",
		Short: "List device commands",
		Long:  "List devices in the Tailscale API",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			all, err := cmd.Flags().GetBool("all")

			if err != nil {
				return fmt.Errorf("failed to get all flag: %w", err)
			}

			var devices []tsapi.Device

			if all {
				devices, err = client.Devices().ListWithAllFields(cmd.Context())
				if err != nil {
					return fmt.Errorf("failed to list devices with all fields: %w", err)
				}
			} else {

				devices, err = client.Devices().List(cmd.Context())
				if err != nil {
					return fmt.Errorf("failed to list devices: %w", err)
				}
			}

			out, err := json.MarshalIndent(devices, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal devices into JSON: %w", err)
			}
			fmt.Fprintln(os.Stdout, string(out))
			return nil

		},
	}

	command.Flags().Bool("all", false, "Display all fields in result.")

	return command
}
