package logs

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/list/logs/config"
	"github.com/jaxxstorm/tscli/cmd/tscli/list/logs/network"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "logs",
		Short: "Get logs for the tailnet",
		Long:  "Commands to retrieve configuration and network audit logs from the Tailscale API.",
	}

	command.AddCommand(config.Command())
	command.AddCommand(network.Command())

	return command
}
