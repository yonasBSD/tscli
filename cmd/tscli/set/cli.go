package set

import (

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "set",
		Short: "Set commands",
		Long:  "Commands that set information on the Tailscale API",
	}



	return command
}
