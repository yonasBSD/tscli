package set

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/set/contact"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/device"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/dns"
	postureintegration "github.com/jaxxstorm/tscli/cmd/tscli/set/integration"

	"github.com/jaxxstorm/tscli/cmd/tscli/set/settings"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/user"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/policy"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "set",
		Short: "Set commands",
		Long:  "Commands that set information on the Tailscale API",
	}

	command.AddCommand(device.Command())
	command.AddCommand(user.Command())
	command.AddCommand(settings.Command())
	command.AddCommand(contact.Command())
	command.AddCommand(postureintegration.Command())
	command.AddCommand(dns.Command())
	command.AddCommand(policy.Command())
	return command
}
