// cmd/tscli/devices/routes/command.go
//
// `tscli devices routes --device <id> --route 10.0.0.0/24 --route 192.168.1.0/24`
// Replaces the enabled subnet-routes list on the device.

package routes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var (
		deviceID string
		routes   []string
	)

	cmd := &cobra.Command{
		Use:   "routes",
		Short: "Set enabled subnet routes for a device",
		Long: `Replace the current list of **enabled** subnet routes for a device.

Examples

  # Replace with two CIDRs
  tscli devices routes --device node-abc123 \
      --route 10.0.0.0/24 --route 192.168.1.0/24
`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if len(routes) == 0 {
				return errors.New("at least one --route must be supplied")
			}
			for _, r := range routes {
				if _, _, err := net.ParseCIDR(r); err != nil {
					return fmt.Errorf("invalid CIDR %q: %w", r, err)
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			if err := client.Devices().SetSubnetRoutes(
				context.Background(),
				deviceID,
				routes,
			); err != nil {
				return fmt.Errorf("failed to set subnet routes: %w", err)
			}

			resp := map[string]any{
				"result": "subnet routes updated",
				"device": deviceID,
				"routes": routes,
			}
			out, _ := json.MarshalIndent(resp, "", "  ")
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}

	// ------------- flags ----------------
	cmd.Flags().StringVar(
		&deviceID,
		"device",
		"",
		"Device ID to modify (nodeId or numeric id)",
	)
	cmd.Flags().StringSliceVar(
		&routes,
		"route",
		nil,
		"CIDR block to enable (repeatable). Example: --route 10.0.0.0/24 --route 192.168.1.0/24",
	)

	_ = cmd.MarkFlagRequired("device")
	_ = cmd.MarkFlagRequired("route")

	return cmd
}
