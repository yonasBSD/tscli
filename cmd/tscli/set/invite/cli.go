// cmd/tscli/set/invite/cli.go

package invite

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/set/invite/device"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "invite",
		Short: "Set invite commands",
		Long:  "Commands that set invite information on the Tailscale API",
	}

	command.AddCommand(device.Command())

	return command
}
