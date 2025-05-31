// cmd/tscli/list/posture-integrations/cli.go
package postureintegration

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

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "posture-integrations",
		Short: "List posture integrations",
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

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
			format := viper.GetString("format")
			output.Print(format, out)
			return nil
		},
	}
}
