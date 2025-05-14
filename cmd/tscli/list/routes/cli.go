// cmd/tscli/list/routes/cli.go

package routes

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "routes",
		Short: "List route commands",
		Long:  "List routes on a device in the Tailscale API",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			deviceID, err := cmd.Flags().GetString("device")
			if err != nil {
				return fmt.Errorf("failed to get device flag: %w", err)
			}

			routes, err := client.Devices().SubnetRoutes(cmd.Context(), deviceID)

			if err != nil {
				return fmt.Errorf("failed to list routes for device %s: %w", deviceID, err)
			}

			out, err := json.MarshalIndent(routes, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal routes into JSON: %w", err)
			}
			fmt.Fprintln(os.Stdout, string(out))
			return nil

		},
	}

	command.Flags().String("device", "", "Device ID to get.")
	_ = command.MarkFlagRequired("device")

	return command
}
