package version

import (
	"fmt"

	"github.com/jaxxstorm/tscli/pkg/version"
	"github.com/jaxxstorm/vers"
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "version",
		Short: "Get the current version",
		Long:  `Get the current version of tscli`,
		RunE: func(*cobra.Command, []string) error {

			v := version.Version
			// If we haven't set a version with linker flags, calculate from git
			if v == "" {
				repo, err := vers.OpenRepository(".")
				if err != nil {
					return fmt.Errorf("error opening repository: %w", err)
				}

				opts := vers.Options{
					Repository: repo,
					Commitish:  "HEAD",
				}

				versions, err := vers.Calculate(opts)
				if err != nil {
					return fmt.Errorf("error calculating version: %w", err)
				}
				v = versions.SemVer
			}
			fmt.Println(v)
			return nil
		},
	}
	return command
}
