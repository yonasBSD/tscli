package user

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/set/user/access"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/user/role"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "user",
		Short: "User commands",
		Long:  "Commands that set User properties on the Tailscale API",
	}

	command.AddCommand(access.Command())
	command.AddCommand(role.Command())

	return command
}
