package dns

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/get/dns/nameservers"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "dns",
		Short: "Get DNS information",
		Long:  "Get information about DNS settings.",
	}

	command.AddCommand(nameservers.Command())

	return command
}
