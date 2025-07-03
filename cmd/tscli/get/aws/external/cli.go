package external
// cmd/tscli/get/aws/external/cli.go
//
// `tscli get aws external-id`
// Get or create AWS external ID for log streaming integration
package external

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
	cmd := &cobra.Command{
		Use:   "external-id",
		Short: "Get AWS external ID for log streaming",
		Long: `Get or create an AWS external ID for use in IAM role trust policies.
This external ID is required when setting up AWS log streaming integration.`,
		
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			// Try the most likely endpoint for AWS external ID
			var response map[string]interface{}
			if _, err := tscli.Do(
				context.Background(),
				client,
				http.MethodGet,
				"/tailnet/{tailnet}/logging/aws/external-id",
				nil,
				&response,
			); err != nil {
				// If that doesn't work, try a POST to create it
				if _, err := tscli.Do(
					context.Background(),
					client,
					http.MethodPost,
					"/tailnet/{tailnet}/logging/aws/external-id",
					nil,
					&response,
				); err != nil {
					return fmt.Errorf("failed to get or create AWS external ID: %w", err)
				}
			}

			out, err := json.MarshalIndent(response, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal response: %w", err)
			}
			
			outputType := viper.GetString("output")
			output.Print(outputType, out)
			return nil
		},
	}

	return cmd
}
