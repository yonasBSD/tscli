// cmd/tscli/delete/invite/user/cli.go
package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jaxxstorm/tscli/pkg/output"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Command registers:  tscli delete invite user --id <invite-id>
func Command() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "user",
		Short: "Delete a user invite",
		RunE: func(cmd *cobra.Command, _ []string) error {

			client, err := tscli.New()
			if err != nil {
				return err
			}

			if _, err := tscli.Do(
				context.Background(),
				client,
				http.MethodDelete,
				fmt.Sprintf("/user-invites/%s", id),
				nil,
				nil,
			); err != nil {
				return fmt.Errorf("delete invite failed: %w", err)
			}

			resp := map[string]string{
				"result": fmt.Sprintf("user invite %s: deleted", id),
			}

			out, _ := json.MarshalIndent(resp, "", "  ")
			format := viper.GetString("format")
			output.Print(format, out)

			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Invite ID to delete (required)")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}
