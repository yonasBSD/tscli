package device

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/device/authorize"
	"github.com/jaxxstorm/tscli/cmd/tscli/device/get"
	"github.com/jaxxstorm/tscli/cmd/tscli/device/list"
	"github.com/jaxxstorm/tscli/cmd/tscli/device/expire"
	"github.com/jaxxstorm/tscli/cmd/tscli/device/delete"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "device",
		Short: "Device commands",
		Long:  "Commands that operate on device",
	}

	command.AddCommand(list.Command())
	command.AddCommand(get.Command())
	command.AddCommand(authorize.Command())
	command.AddCommand(expire.Command())
	command.AddCommand(delete.Command())

	return command
}
