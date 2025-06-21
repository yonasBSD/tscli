// cmd/tscli/set/name/cli.go

package name

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jaxxstorm/tscli/pkg/output"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Command() *cobra.Command {
	var (
		deviceID string
		newName  string
	)

	cmd := &cobra.Command{
		Use:   "name",
		Short: "Set a device name",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			if err := client.Devices().SetName(context.Background(), deviceID, newName); err != nil {
				return fmt.Errorf("failed to set device name: %w", err)
			}

			out, _ := json.MarshalIndent(map[string]string{
				"result": fmt.Sprintf("device %s name set to %s", deviceID, newName),
			}, "", "  ")
			outputType := viper.GetString("output")
			output.Print(outputType, out)
			return nil
		},
	}

	cmd.Flags().StringVar(&deviceID, "device", "", "Device ID to rename")
	cmd.Flags().StringVar(&newName, "name", "", "New device hostname")
	_ = cmd.MarkFlagRequired("device")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}
