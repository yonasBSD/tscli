package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type PrettyPrinter struct{}

func (PrettyPrinter) Print(raw []byte) error {
	var payload any
	if err := json.NewDecoder(bytes.NewReader(raw)).Decode(&payload); err != nil {
		return err
	}

	recs := normalise(payload)
	if len(recs) == 0 {
		fmt.Println("no records found")
		return nil
	}

	for i, rec := range recs {
		if err := renderMap(rec, 0); err != nil {
			return err
		}
		if i != len(recs)-1 {
			fmt.Println(div.Render(strings.Repeat("─", termWidth())))
		}
	}
	return nil
}

var (
	keyStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7D56F4"))
	valStyle  = lipgloss.NewStyle()
	div       = lipgloss.NewStyle().Foreground(lipgloss.Color("#4B4B4B"))
	boolTrue  = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).Render("✔")
	boolFalse = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5F87")).Render("✖")
	none      = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666")).Render("—")

	maxScalar      = 80 // truncate very long strings
	maxInlineArray = 10
	maxInlineKeys  = 6
	padding        = 2
)

func termWidth() int {
	if w, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil && w > 0 {
		return w
	}
	return 100
}

func renderMap(m map[string]any, indent int) error {
	keys := mapsKeys(m)
	slices.Sort(keys)

	longest := 0
	for _, k := range keys {
		if len(k) > longest {
			longest = len(k)
		}
	}
	keyWidth := longest + 2

	for _, k := range keys {
		keyCell := keyStyle.Width(keyWidth).Render(k + ":")

		rawVal := fmtPretty(m[k], indent+keyWidth+padding)

		if strings.Contains(rawVal, "\n") {
			// multiline value: key on its own line
			fmt.Println(strings.Repeat(" ", indent) + keyCell)
			for _, ln := range strings.Split(rawVal, "\n") {
				if ln == "" {
					continue
				}
				fmt.Println(strings.Repeat(" ", indent+keyWidth) + ln)
			}
			continue
		}

		// single-line value
		valCell := wrap(valStyle, rawVal, termWidth()-keyWidth-indent-padding)
		fmt.Println(strings.Repeat(" ", indent) +
			lipgloss.JoinHorizontal(lipgloss.Top, keyCell, valCell))
	}
	return nil
}

func fmtPretty(v any, _ int) string {
	switch x := v.(type) {

	case nil:
		return none

	case string:
		if x == "" {
			return none
		}
		if len(x) > maxScalar {
			return x[:maxScalar-1] + "…"
		}
		return x

	case bool:
		if x {
			return boolTrue
		}
		return boolFalse

	case json.Number:
		return x.String()

	case []any:
		if len(x) == 0 {
			return "[]"
		}
		/* inline small scalar arrays */
		var parts []string
		for _, el := range x {
			switch el.(type) {
			case string, json.Number, bool:
				parts = append(parts, fmtPretty(el, 0))
			default:
				return fmt.Sprintf("[%d items]", len(x))
			}
			if len(parts) > maxInlineArray {
				return fmt.Sprintf("[%d items]", len(x))
			}
		}
		inline := "[" + strings.Join(parts, ", ") + "]"
		if len(inline) > maxScalar {
			return fmt.Sprintf("[%d items]", len(x))
		}
		return inline

	case map[string]any:
		if len(x) == 0 {
			return "{}"
		}
		/* inline very small maps */
		if len(x) <= maxInlineKeys {
			ks := mapsKeys(x)
			slices.Sort(ks)
			var pieces []string
			total := 2
			for _, k := range ks {
				p := k + ": " + fmtPretty(x[k], 0)
				total += len(p) + 2
				if total > maxScalar {
					goto multiline
				}
				pieces = append(pieces, p)
			}
			return "{" + strings.Join(pieces, ", ") + "}"
		}

	multiline:
		var b strings.Builder
		for _, k := range mapsKeys(x) {
			b.WriteString(k)
			b.WriteString(": ")
			b.WriteString(fmtPretty(x[k], 0))
			b.WriteByte('\n')
		}
		return b.String()

	default:
		j, _ := json.Marshal(x)
		s := string(j)
		if len(s) > maxScalar {
			return s[:maxScalar-1] + "…"
		}
		return s
	}
}

func wrap(st lipgloss.Style, s string, width int) string {
	return st.MaxWidth(width).Render(s)
}

func mapsKeys(m map[string]any) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

func normalise(v any) []map[string]any {
	switch x := v.(type) {
	case []any:
		out := make([]map[string]any, 0, len(x))
		for _, el := range x {
			if m, ok := el.(map[string]any); ok {
				out = append(out, m)
			}
		}
		return out
	case map[string]any:
		return []map[string]any{x}
	default:
		return nil
	}
}
