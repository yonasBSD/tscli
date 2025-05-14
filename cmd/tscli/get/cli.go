package get

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/get/device"
	"github.com/jaxxstorm/tscli/cmd/tscli/get/posture"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "get",
		Short: "Get commands",
		Long:  "Commands that get information from the Tailscale API",
	}

	command.AddCommand(device.Command())
	command.AddCommand(posture.Command())

	return command
}
