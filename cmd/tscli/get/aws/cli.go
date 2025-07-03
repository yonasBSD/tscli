package aws
// cmd/tscli/get/aws/cli.go
//
// `tscli get aws external-id`
// Get or create AWS external ID for log streaming integration
package aws

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/get/aws/external"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "aws",
		Short: "Get AWS integration information",
		Long:  "Commands to retrieve AWS integration information from the Tailscale API.",
	}

	command.AddCommand(external.Command())

	return command
}
