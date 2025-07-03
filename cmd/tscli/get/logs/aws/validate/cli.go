// cmd/tscli/get/logs/aws/validate/cli.go
//
// `tscli get logs aws validate --external-id <id> --role-arn <arn>`
// Validate AWS trust policy for log streaming integration
package validate

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
	var (
		externalID string
		roleArn    string
	)

	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate AWS trust policy for log streaming",
		Long: `Validate the AWS IAM role trust policy for log streaming integration.
This validates that the external ID and role ARN are correctly configured.`,

		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return fmt.Errorf("failed to create client: %w", err)
			}

			requestBody := map[string]interface{}{
				"roleArn": roleArn,
			}

			var response map[string]interface{}
			if _, err := tscli.Do(
				context.Background(),
				client,
				http.MethodPost,
				fmt.Sprintf("/tailnet/{tailnet}/aws-external-id/%s/validate-aws-trust-policy", externalID),
				requestBody,
				&response,
			); err != nil {
				return fmt.Errorf("failed to validate AWS trust policy: %w", err)
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

	cmd.Flags().StringVar(&externalID, "external-id", "", "AWS external ID (required)")
	cmd.Flags().StringVar(&roleArn, "role-arn", "", "AWS IAM role ARN (required)")

	_ = cmd.MarkFlagRequired("external-id")
	_ = cmd.MarkFlagRequired("role-arn")

	return cmd
}
