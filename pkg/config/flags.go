// pkg/config/flags.go
package config

import (
	"fmt"
	"io"
	"sort"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	yellow = "\033[33;1m" // bright-yellow, bold
	reset  = "\033[0m"
	warn   = "⚠️ "
)

// ChangedMap gathers the flag names the user explicitly set.
func ChangedMap(cmd *cobra.Command) map[string]struct{} {
	m := make(map[string]struct{})

	appendChanged := func(fs *pflag.FlagSet) {
		if fs == nil {
			return
		}
		fs.Visit(func(f *pflag.Flag) { m[f.Name] = struct{}{} })
	}
	appendChanged(cmd.Flags())
	appendChanged(cmd.InheritedFlags())
	return m
}

// WarnIfOverrides prints a coloured warning if any flags override config/env.
func WarnIfOverrides(w io.Writer, cmd *cobra.Command) {
	changed := ChangedMap(cmd)
	if len(changed) == 0 {
		return
	}

	keys := make([]string, 0, len(changed))
	for k := range changed {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Fprintf(
		w,
		"%s%s WARNING%s overriding config values with flags: %v\n",
		yellow, warn, reset, keys,
	)
}
