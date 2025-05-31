// cmd/tscli/get/posture/cli.go
//
// `tscli get posture --device <id>`
// Fetch the custom posture-attribute map for a device.

package posture

import (
	"encoding/json"
	"fmt"

	"github.com/jaxxstorm/tscli/pkg/output"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Command() *cobra.Command {
	var deviceID string

	cmd := &cobra.Command{
		Use:   "posture",
		Short: "Get posture attributes for a device",
		Long: `Return all custom posture attributes currently set on a device.

Example

  tscli get posture --device node-abcdef123456
`,
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

			attrs, err := client.Devices().GetPostureAttributes(cmd.Context(), deviceID)
			if err != nil {
				return fmt.Errorf("failed to get posture attributes: %w", err)
			}

			out, err := json.MarshalIndent(attrs, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal JSON: %w", err)
			}
			format := viper.GetString("format")
			output.Print(format, out)
			return nil
		},
	}

	// ---------------- flags ----------------
	cmd.Flags().StringVar(
		&deviceID,
		"device",
		"",
		`Device ID to query (nodeId "node-abc123" or numeric id). Example: --device node-abcdef123456`,
	)
	_ = cmd.MarkFlagRequired("device")

	return cmd
}
