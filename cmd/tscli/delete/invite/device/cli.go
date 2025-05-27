// cmd/tscli/delete/invite/device/cli.go
package device

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
)

// Command registers:  tscli delete invite device --id <invite-id>
func Command() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "device",
		Short: "Delete a device share invite",
		RunE: func(cmd *cobra.Command, _ []string) error {

			client, err := tscli.New()
			if err != nil {
				return err
			}

			if _, err := tscli.Do(
				context.Background(),
				client,
				http.MethodDelete,
				fmt.Sprintf("/device-invites/%s", id),
				nil,
				nil,
			); err != nil {
				return fmt.Errorf("delete invite failed: %w", err)
			}

			fmt.Fprintf(os.Stdout, `{"result":"device invite deleted","id":"%s"}`+"\n", id)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Invite ID to delete (required)")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}
