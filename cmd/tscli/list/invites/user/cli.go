// cmd/tscli/list/invites/user/cli.go
package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/jaxxstorm/tscli/pkg/tscli"
	"github.com/spf13/cobra"
)

var validState = map[string]struct{}{
	"pending":  {},
	"accepted": {},
	"all":      {},
}

func Command() *cobra.Command {
	var state string

	cmd := &cobra.Command{
		Use:   "user",
		Short: "List user invites in the tailnet",
		PersistentPreRunE: func(*cobra.Command, []string) error {
			if state != "" {
				if _, ok := validState[strings.ToLower(state)]; !ok {
					return fmt.Errorf("invalid --state %q (pending|accepted|all)", state)
				}
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, _ []string) error {
			client, err := tscli.New()
			if err != nil {
				return err
			}

			path := "/tailnet/{tailnet}/user-invites"
			if s := strings.ToLower(state); s != "" && s != "all" {
				q := url.Values{}
				q.Set("state", s)
				path += "?" + q.Encode()
			}

			var resp json.RawMessage
			if _, err := tscli.Do(
				context.Background(),
				client,
				http.MethodGet,
				path,
				nil,
				&resp,
			); err != nil {
				return err
			}

			pretty, _ := json.MarshalIndent(resp, "", "  ")
			fmt.Fprintln(os.Stdout, string(pretty))
			return nil
		},
	}

	cmd.Flags().StringVar(&state, "state", "",
		`Filter by state: pending | accepted | all`)
	return cmd
}
