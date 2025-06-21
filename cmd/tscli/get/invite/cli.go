package invite

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/get/invite/device"
	"github.com/jaxxstorm/tscli/cmd/tscli/get/invite/user"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "invite",
		Short: "Get invite commands",
		Long:  "Commands that get invites in the Tailscale API",
	}

	command.AddCommand(user.Command())
	command.AddCommand(device.Command())

	return command
}
