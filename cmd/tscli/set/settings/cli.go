// cmd/tscli/set/settings/cli.go
package settings

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	tsapi "tailscale.com/client/tailscale/v2"
)

var validJoin = map[string]struct{}{
	"none": {}, "admin": {}, "member": {},
}

func Command() *cobra.Command {
	var (
		devAppr, devAuto, usrAppr,
		netLog, regRoute, postureID bool
		keyDays  int
		joinRole string
	)

	cmd := &cobra.Command{
		Use:   "settings",
		Short: "Update tailnet settings",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			f := cmd.Flags()

			// require at least one flag
			changed := 0
			f.Visit(func(_ *pflag.Flag) { changed++ })
			if changed == 0 {
				return fmt.Errorf("at least one setting flag must be provided")
			}

			if f.Lookup("users-role-join").Changed {
				joinRole = strings.ToLower(joinRole)
				if _, ok := validJoin[joinRole]; !ok {
					return fmt.Errorf("invalid --users-role-join: %s (none|admin|member)", joinRole)
				}
			}
			if f.Lookup("devices-key-duration").Changed {
				if keyDays < 1 || keyDays > 180 {
					return fmt.Errorf("--devices-key-duration must be 1-180")
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

			req := tsapi.UpdateTailnetSettingsRequest{}
			f := cmd.Flags()

			if f.Lookup("devices-approval").Changed {
				req.DevicesApprovalOn = tsapi.PointerTo(devAppr)
			}
			if f.Lookup("devices-auto-updates").Changed {
				req.DevicesAutoUpdatesOn = tsapi.PointerTo(devAuto)
			}
			if f.Lookup("devices-key-duration").Changed {
				req.DevicesKeyDurationDays = tsapi.PointerTo(keyDays)
			}
			if f.Lookup("users-approval").Changed {
				req.UsersApprovalOn = tsapi.PointerTo(usrAppr)
			}
			if f.Lookup("users-role-join").Changed {
				role := tsapi.RoleAllowedToJoinExternalTailnets(joinRole)
				req.UsersRoleAllowedToJoinExternalTailnets = tsapi.PointerTo(role)
			}
			if f.Lookup("network-flow-logging").Changed {
				req.NetworkFlowLoggingOn = tsapi.PointerTo(netLog)
			}
			if f.Lookup("regional-routing").Changed {
				req.RegionalRoutingOn = tsapi.PointerTo(regRoute)
			}
			if f.Lookup("posture-identity-collection").Changed {
				req.PostureIdentityCollectionOn = tsapi.PointerTo(postureID)
			}

			if err := client.TailnetSettings().Update(context.Background(), req); err != nil {
				return fmt.Errorf("update failed: %w", err)
			}

			out, _ := json.MarshalIndent(req, "", "  ")
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}

	cmd.Flags().BoolVar(&devAppr, "devices-approval", false, "Enable/disable device approval")
	cmd.Flags().BoolVar(&devAuto, "devices-auto-updates", false, "Enable/disable device auto-updates")
	cmd.Flags().BoolVar(&usrAppr, "users-approval", false, "Enable/disable user approval")
	cmd.Flags().BoolVar(&netLog, "network-flow-logging", false, "Enable/disable network-flow logging")
	cmd.Flags().BoolVar(&regRoute, "regional-routing", false, "Enable/disable regional routing")
	cmd.Flags().BoolVar(&postureID, "posture-identity-collection", false, "Enable/disable posture identity collection")

	cmd.Flags().IntVar(&keyDays, "devices-key-duration", 0, "Device key expiry (1-180 days)")
	cmd.Flags().StringVar(&joinRole, "users-role-join", "", "Role allowed to join external tailnets (none|admin|member)")

	return cmd
}
