package devices

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/devices/authorize"
	"github.com/jaxxstorm/tscli/cmd/tscli/devices/get"
	"github.com/jaxxstorm/tscli/cmd/tscli/devices/list"
	
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "devices",
		Short: "Device commands",
		Long:  "Commands that operate on devices",
	}

	command.AddCommand(list.Command())
	command.AddCommand(get.Command())
	command.AddCommand(authorize.Command())

	return command
}