// cmd/tscli/list/posture-integrations/cli.go
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
	return &cobra.Command{
		Use:   "posture-integrations",
		Short: "List posture integrations",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

			// generic container: each integration is decoded as map[string]any
			var resp struct {
				Integrations []map[string]any `json:"integrations"`
			}

			if _, err := tscli.Do(
				context.Background(),
				client,
				http.MethodGet,
				"/tailnet/{tailnet}/posture/integrations",
				nil,
				&resp,
			); err != nil {
				return fmt.Errorf("failed to list posture integrations: %w", err)
			}

			out, _ := json.MarshalIndent(resp, "", "  ")
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}
}
