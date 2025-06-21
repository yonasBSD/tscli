// cmd/tscli/get/user/cli.go
package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jaxxstorm/tscli/pkg/output"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tsapi "tailscale.com/client/tailscale/v2"
)

func Command() *cobra.Command {
	var userID string

	cmd := &cobra.Command{
		Use:   "user",
		Short: "Get a tailnet user",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			var u *tsapi.User
			u, err = client.Users().Get(context.Background(), userID)
			if err != nil {
				return fmt.Errorf("failed to get user %s: %w", userID, err)
			}

			out, err := json.MarshalIndent(u, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal user: %w", err)
			}
			outputType := viper.GetString("output")
			output.Print(outputType, out)
			return nil
		},
	}

	cmd.Flags().StringVar(&userID, "user", "", "User ID (email or UID) to retrieve")
	_ = cmd.MarkFlagRequired("user")

	return cmd
}
