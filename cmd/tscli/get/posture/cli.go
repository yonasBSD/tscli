package posture

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "posture",
		Short: "Get posture commands",
		Long:  "Get commands that return the posture of a device",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			deviceID, err := cmd.Flags().GetString("device")
			if err != nil {
				return fmt.Errorf("failed to get device flag: %w", err)
			}

			attributes, err := client.Devices().GetPostureAttributes(cmd.Context(), deviceID)
			if err != nil {
				return fmt.Errorf("failed to list devices with all fields: %w", err)
			}

			out, err := json.MarshalIndent(attributes, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal attributes into JSON: %w", err)
			}
			fmt.Fprintln(os.Stdout, string(out))
			return nil

		},
	}

	command.Flags().String("device", "", "Device ID to get.")
	_ = command.MarkFlagRequired("device")

	return command
}
