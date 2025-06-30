// cmd/tscli/create/token/cli.go
//
// Exchange an OAuth client-id/secret for an access-token.
//
//   tscli create token --client-id XXX --client-secret YYY
//
package token

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/jaxxstorm/tscli/pkg/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const endpoint = "https://api.tailscale.com/api/v2/oauth/token"

func Command() *cobra.Command {
	var (
		clientID     string
		clientSecret string
	)

	cmd := &cobra.Command{
		Use:   "token",
		Short: "Create an OAuth access-token from a client_id / client_secret pair",
		Long:  "POSTs to /api/v2/oauth/token. No API-key auth is required.",
		Example: `  tscli create token \
      --client-id     "$OAUTH_CLIENT_ID" \
      --client-secret "$OAUTH_CLIENT_SECRET"`,
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			if clientID == "" || clientSecret == "" {
				return errors.New("--client-id and --client-secret are both required")
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, _ []string) error {
			/* ------------------------------------------------------ */
			form := url.Values{}
			form.Set("client_id", clientID)
			form.Set("client_secret", clientSecret)

			req, _ := http.NewRequest(http.MethodPost, endpoint,
				bytes.NewBufferString(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("User-Agent", "tscli")

			if viper.GetBool("debug") {
				if dump, err := httputil.DumpRequestOut(req, true); err == nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "\n--- REQUEST ---\n%s\n", dump)
				}
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			if viper.GetBool("debug") {
				if dump, err := httputil.DumpResponse(resp, true); err == nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "\n--- RESPONSE ---\n%s\n", dump)
				}
			}

			body, _ := io.ReadAll(resp.Body)

			/* ------------------------------------------------------ */
			if resp.StatusCode/100 != 2 { // non-2xx
				_ = output.Print(viper.GetString("format"), body)
				return fmt.Errorf("tailscale API returned %s", resp.Status)
			}

			return output.Print(viper.GetString("format"), body)
		},
	}

	cmd.Flags().StringVar(&clientID, "client-id", "", "OAuth client ID (required)")
	cmd.Flags().StringVar(&clientSecret, "client-secret", "", "OAuth client secret (required)")
	_ = cmd.MarkFlagRequired("client-id")
	_ = cmd.MarkFlagRequired("client-secret")

	return cmd
}
