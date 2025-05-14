package device

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/device/authorize"
	"github.com/jaxxstorm/tscli/cmd/tscli/device/expire"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "device",
		Short: "Device commands",
		Long:  "Commands that operate on device",
	}

	command.AddCommand(authorize.Command())
	command.AddCommand(expire.Command())

	return command
}
