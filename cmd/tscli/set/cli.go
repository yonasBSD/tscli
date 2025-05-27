package set

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/set/authorization"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/contact"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/dns"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/expiry"
	postureintegration "github.com/jaxxstorm/tscli/cmd/tscli/set/integration"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/ip"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/posture"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/routes"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/settings"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/tags"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/user"
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
	command.AddCommand(user.Command())
	command.AddCommand(settings.Command())
	command.AddCommand(contact.Command())
	command.AddCommand(postureintegration.Command())
	command.AddCommand(authorization.Command())
	command.AddCommand(expiry.Command())
	command.AddCommand(dns.Command())
	return command
}
