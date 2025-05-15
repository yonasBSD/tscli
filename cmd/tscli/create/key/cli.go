// cmd/tscli/create/key/cli.go
package key

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
)

var keyTypes = map[string]struct{}{
	"authkey":     {},
	"oauthclient": {},
}

func Command() *cobra.Command {
	var (
		kType       string
		description string
		expiry      int      // seconds (authkey)
		scopes      []string // oauthclient
		tags        []string // oauthclient
	)

	cmd := &cobra.Command{
		Use:   "key",
		Short: "Create an auth-key or OAuth client",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			kType = strings.ToLower(kType)
			if kType == "" {
				kType = "authkey"
			}
			if _, ok := keyTypes[kType]; !ok {
				return fmt.Errorf("invalid --type: %s (authkey|oauthclient)", kType)
			}

			if kType == "oauthclient" {
				if len(scopes) == 0 {
					return fmt.Errorf("--scope is required for oauthclient")
				}
				if requiresTags(scopes) && len(tags) == 0 {
					return fmt.Errorf("--tags required when scope includes devices:core or auth_keys")
				}
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

			// Build body
			body := map[string]any{
				"keyType": strings.ToLower(kType),
			}
			if description != "" {
				body["description"] = description
			}

			if kType == "authkey" {
				body["capabilities"] = map[string]any{
					"devices": map[string]any{},
				}
				if cmd.Flags().Lookup("expiry").Changed {
					body["expirySeconds"] = expiry
				}
			} else { // oauthclient
				body["scopes"] = scopes
				if len(tags) > 0 {
					body["tags"] = tags
				}
			}

			var resp map[string]any
			if _, err := tscli.Do(
				context.Background(),
				client,
				http.MethodPost,
				"/tailnet/{tailnet}/keys",
				body,
				&resp,
			); err != nil {
				return fmt.Errorf("key creation failed: %w", err)
			}

			out, _ := json.MarshalIndent(resp, "", "  ")
			fmt.Fprintln(os.Stdout, string(out))
			return nil
		},
	}

	cmd.Flags().StringVar(&kType, "type", "authkey", "Key type: authkey|oauthclient")
	cmd.Flags().StringVar(&description, "description", "", "Short description (≤50 chars)")
	cmd.Flags().IntVar(&expiry, "expiry", 0, "Expiry in seconds (authkey only)")
	cmd.Flags().StringSliceVar(&scopes, "scope", nil, "OAuth scopes (repeatable) (oauthclient)")
	cmd.Flags().StringSliceVar(&tags, "tags", nil, "Allowed tags (repeatable) (oauthclient)")

	return cmd
}

func requiresTags(scopes []string) bool {
	for _, s := range scopes {
		if s == "devices:core" || s == "auth_keys" {
			return true
		}
	}
	return false
}
