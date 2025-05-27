// cmd/tscli/get/device/cli.go
//
// Fetch details for a single device.
//
//	# by node-ID
//	tscli get device --device node-abcdef123456
//
//	# by Tailscale IP
//	tscli get device --ip 100.64.0.12
//
//	# by hostname (case-insensitive)
//	tscli get device --name db-server
//
//	# include every possible field
//	tscli get device --device node-abcdef123456 --all
package device

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"slices"
	"strings"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	tsapi "tailscale.com/client/tailscale/v2"
)

var newClient = tscli.New

func Command() *cobra.Command {
	var (
		showAll bool

		deviceID string
		deviceIP string
		devName  string
	)

	cmd := &cobra.Command{
		Use:   "device",
		Short: "Get a device's information",
		Long:  "Return a single device record from the Tailscale API.",
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			chosen := 0
			if deviceID != "" {
				chosen++
			}
			if deviceIP != "" {
				if net.ParseIP(deviceIP) == nil {
					return fmt.Errorf("invalid --ip %q", deviceIP)
				}
				chosen++
			}
			if devName != "" {
				chosen++
			}

			switch chosen {
			case 0:
				return errors.New("one of --device, --ip or --name is required")
			case 1:
				return nil
			default:
				return errors.New("--device, --ip and --name are mutually exclusive")
			}
		},

		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := newClient()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			if deviceID == "" {
				devices, err := client.Devices().List(cmd.Context())
				if err != nil {
					return fmt.Errorf("device lookup failed: %w", err)
				}

				for _, d := range devices {
					switch {
					case deviceIP != "" && slices.Contains(d.Addresses, deviceIP):
						deviceID = d.NodeID
					case devName != "" && strings.EqualFold(d.Hostname, devName):
						deviceID = d.NodeID
					}
					if deviceID != "" {
						break // stop at the first match
					}
				}
				if deviceID == "" {
					return errors.New("no matching device found")
				}
			}

			var dv *tsapi.Device
			if showAll {
				dv, err = client.Devices().GetWithAllFields(cmd.Context(), deviceID)
			} else {
				dv, err = client.Devices().Get(cmd.Context(), deviceID)
			}
			if err != nil {
				return fmt.Errorf("failed to get device %s: %w", deviceID, err)
			}

			out, _ := json.MarshalIndent(dv, "", "  ")
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}

	/* ----------------------- flags -------------------------------- */
	cmd.Flags().BoolVar(&showAll, "all", false,
		"Include advanced fields such as ClientConnectivity, AdvertisedRoutes, and EnabledRoutes.")

	cmd.Flags().StringVar(&deviceID, "device", "",
		`Device identifier (preferred nodeId "node-abc123" or legacy numeric id).`)

	cmd.Flags().StringVar(&deviceIP, "ip", "",
		"Match by the device’s Tailscale IPv4 (or IPv6) address.")

	cmd.Flags().StringVar(&devName, "name", "",
		"Match by device hostname (case-insensitive).")

	return cmd
}


