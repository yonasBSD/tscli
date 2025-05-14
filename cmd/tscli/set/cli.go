package set

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/set/ip"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/posture"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/routes"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/tags"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "set",
		Short: "Set commands",
		Long:  "Commands that set information on the Tailscale API",
	}

	command.AddCommand(posture.Command())
	command.AddCommand(routes.Command())
	command.AddCommand(tags.Command())
	command.AddCommand(ip.Command())

	return command
}
