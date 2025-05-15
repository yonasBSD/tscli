// cmd/tscli/delete/webhook/delete.go
package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var hookID string

	cmd := &cobra.Command{
		Use:   "webhook",
		Short: "Deletea webhook",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

			if err := client.Webhooks().Delete(context.Background(), hookID); err != nil {
				return err
			}

			out, _ := json.MarshalIndent(map[string]string{
				"result": fmt.Sprintf("webhook %s deleted", hookID),
			}, "", "  ")
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}

	cmd.Flags().StringVar(&hookID, "id", "", "Webhook ID to delete")
	_ = cmd.MarkFlagRequired("id")
	return cmd
}
