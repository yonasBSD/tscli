package delete

import (
	"github.com/jaxxstorm/tscli/cmd/tscli/delete/device"
	postureintegration "github.com/jaxxstorm/tscli/cmd/tscli/delete/integration"
	"github.com/jaxxstorm/tscli/cmd/tscli/delete/invite"
	"github.com/jaxxstorm/tscli/cmd/tscli/delete/key"
	"github.com/jaxxstorm/tscli/cmd/tscli/delete/posture"
	"github.com/jaxxstorm/tscli/cmd/tscli/delete/user"
	"github.com/jaxxstorm/tscli/cmd/tscli/delete/webhook"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "delete",
		Short: "Delete commands",
		Long:  "Commands that delete from the Tailscale API",
	}

	command.AddCommand(device.Command())
	command.AddCommand(user.Command())
	command.AddCommand(key.Command())
	command.AddCommand(webhook.Command())
	command.AddCommand(posture.Command())
	command.AddCommand(invite.Command())
	command.AddCommand(postureintegration.Command())

	return command
}
