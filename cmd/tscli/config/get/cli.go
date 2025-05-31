package get

import (
	"fmt"

	"github.com/jaxxstorm/tscli/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "get <key>",
		Short: "Print a single config value",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			config.WarnIfOverrides(cmd.ErrOrStderr(), cmd)

			key := args[0]

			if !viper.IsSet(key) {
				return fmt.Errorf("%q not set", key)
			}
			// Warn if that key was supplied as a flag
			if _, ok := config.ChangedMap(cmd.Root())[key]; ok {
				fmt.Fprintf(cmd.ErrOrStderr(),
					"⚠️  %q is being taken from the command-line, not the config file\n", key)
			}

			fmt.Println(viper.Get(key))
			return nil
		},
	}
}
