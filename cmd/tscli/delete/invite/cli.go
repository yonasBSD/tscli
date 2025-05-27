package invite

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/delete/invite/device"
	"github.com/jaxxstorm/tscli/cmd/tscli/delete/invite/user"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "invite",
		Short: "Delete invite commands",
		Long:  "Commands that delete invites in the Tailscale API",
	}

	command.AddCommand(user.Command())
	command.AddCommand(device.Command())

	return command
}
