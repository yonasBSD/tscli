package dns

import (
	"github.com/spf13/cobra"
	"github.com/jaxxstorm/tscli/cmd/tscli/set/dns/nameservers"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "dns",
		Short: "Set DNS information",
		Long:  "Set information about DNS settings.",
	}

	command.AddCommand(nameservers.Command())

	return command
}
