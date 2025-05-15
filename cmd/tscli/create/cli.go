package create

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/create/integration"
	"github.com/jaxxstorm/tscli/cmd/tscli/create/key"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "create",
		Short: "Create commands",
		Long:  "Commands that create in the Tailscale API",
	}

	command.AddCommand(postureintegration.Command())
	command.AddCommand(key.Command())

	return command
}
