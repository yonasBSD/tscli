package invites

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/list/invites/device"
	"github.com/jaxxstorm/tscli/cmd/tscli/list/invites/user"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "invites",
		Short: "Invite commands",
		Long:  "Commands that list devices from the Tailscale API",
	}

	command.AddCommand(user.Command())
	command.AddCommand(device.Command())

	return command
}
