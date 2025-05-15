// cmd/tscli/get/posture-integration/cli.go
//
// `tscli get posture-integration --id <integration-id>`
// Fetch a single device-posture integration by its identifier.
//
package postureintegration

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "posture-integration",
		Short: "Get a posture integration",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

			var integ map[string]any
			if _, err := tscli.Do(
				context.Background(),
				client,
				http.MethodGet,
				"/tailnet/{tailnet}/posture/integrations/"+id,
				nil,
				&integ,
			); err != nil {
				return fmt.Errorf("failed to get posture integration %s: %w", id, err)
			}

			out, _ := json.MarshalIndent(integ, "", "  ")
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Integration identifier")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}
