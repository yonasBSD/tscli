package version

import (
	"fmt"

	"github.com/jaxxstorm/tscli/pkg/version"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "version",
		Short: "Get the current version",
		Long:  `Get the current version of tscli`,
		RunE: func(*cobra.Command, []string) error {
			v := version.GetVersion()
			fmt.Println(v)
			return nil
		},
	}
	return command
}
