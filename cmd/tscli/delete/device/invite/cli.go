// cmd/tscli/delete/device/invite/cli.go
package invite

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

// Command registers:  tscli delete device invite --id <invite-id>
func Command() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "invite",
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

			resp := map[string]string{
				"result": fmt.Sprintf("device invite %s: deleted", id),
			}

			out, _ := json.MarshalIndent(resp, "", "  ")
			outputType := viper.GetString("output")
			output.Print(outputType, out)

			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Device invite ID to delete")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}
