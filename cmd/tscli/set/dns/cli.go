package dns

import (
	"github.com/spf13/cobra"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/dns/nameservers"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/dns/preferences"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/dns/searchpaths"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "dns",
		Short: "Set DNS information",
		Long:  "Set information about DNS settings.",
	}

	command.AddCommand(nameservers.Command())
	command.AddCommand(preferences.Command())
	command.AddCommand(searchpaths.Command())

	return command
}
