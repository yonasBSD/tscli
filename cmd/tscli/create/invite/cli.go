package invite

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/create/invite/device"
	"github.com/jaxxstorm/tscli/cmd/tscli/create/invite/user"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "invite",
		Short: "Create invite commands",
		Long:  "Commands that create invites in the Tailscale API",
	}

	command.AddCommand(user.Command())
	command.AddCommand(device.Command())

	return command
}
