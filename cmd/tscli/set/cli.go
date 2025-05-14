package set

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/set/posture"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "set",
		Short: "Set commands",
		Long:  "Commands that set information on the Tailscale API",
	}

	command.AddCommand(posture.Command())

	return command
}
