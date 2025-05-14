package list

import (

	"github.com/spf13/cobra"
	"github.com/jaxxstorm/tscli/cmd/tscli/list/devices"
	"github.com/jaxxstorm/tscli/cmd/tscli/list/routes"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "list",
		Short: "List commands",
		Long:  "Commands that list information from the Tailscale API",
	}

	command.AddCommand(devices.Command())
	command.AddCommand(routes.Command())



	return command
}
