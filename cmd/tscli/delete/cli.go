package delete

import (

	"github.com/spf13/cobra"
	"github.com/jaxxstorm/tscli/cmd/tscli/delete/device"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "delete",
		Short: "Delete commands",
		Long:  "Commands that delete from the Tailscale API",
	}

	command.AddCommand(device.Command())


	return command
}
