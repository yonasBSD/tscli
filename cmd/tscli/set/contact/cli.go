// cmd/tscli/set/contacts/cli.go
package contact

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jaxxstorm/tscli/pkg/output"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tsapi "tailscale.com/client/tailscale/v2"
)

var validTypes = map[string]tsapi.ContactType{
	"primary":  tsapi.ContactType("primary"),
	"billing":  tsapi.ContactType("billing"),
	"security": tsapi.ContactType("security"),
}

func Command() *cobra.Command {
	var (
		typeStr string
		email   string
	)

	cmd := &cobra.Command{
		Use:   "contact",
		Short: "Update a contact record in the Tailscale API",
		Long:  "Allowed types: primary | billing | security.  Currently only the email field is writable.",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if email == "" {
				return fmt.Errorf("--email is required")
			}
			if _, ok := validTypes[strings.ToLower(typeStr)]; !ok {
				return fmt.Errorf("invalid --type: %s (primary|billing|security)", typeStr)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

			req := tsapi.UpdateContactRequest{
				Email: tsapi.PointerTo(email),
			}

			ctype := validTypes[strings.ToLower(typeStr)]
			if err := client.Contacts().Update(
				context.Background(),
				ctype, // only 2 args: ctx + type + request
				req,
			); err != nil {
				return fmt.Errorf("update failed: %w", err)
			}

			out, _ := json.MarshalIndent(map[string]string{
				"result": "contact updated",
				"type":   strings.ToLower(typeStr),
				"email":  email,
			}, "", "  ")
			outputType := viper.GetString("output")
			output.Print(outputType, out)
			return nil
		},
	}

	cmd.Flags().StringVar(&typeStr, "type", "primary", "Contact type to update (primary|billing|security)")
	cmd.Flags().StringVar(&email, "email", "", "New email address")
	_ = cmd.MarkFlagRequired("email")

	return cmd
}
