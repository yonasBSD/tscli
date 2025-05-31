package time

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

var durRE = regexp.MustCompile(`(\d+)([dhms])`)

// Accepts sequences like “30d12h45m10s”. Supported units: d, h, m, s.
func ParseHumanDuration(s string) (time.Duration, error) {
	matches := durRE.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 || durRE.ReplaceAllString(s, "") != "" {
		return 0, errors.New("bad duration")
	}

	var total time.Duration
	for _, m := range matches {
		val, _ := strconv.Atoi(m[1])
		switch m[2] {
		case "d":
			total += time.Duration(val) * 24 * time.Hour
		case "h":
			total += time.Duration(val) * time.Hour
		case "m":
			total += time.Duration(val) * time.Minute
		case "s":
			total += time.Duration(val) * time.Second
		}
	}
	return total, nil
}

// parseTime parses either RFC3339 or a relative duration like "2h30m" / "10d".
// If empty and allowNow==true it returns time.Now().UTC().
func ParseTime(in string, allowNow bool) (time.Time, error) {
	if in == "" {
		if allowNow {
			return time.Now().UTC(), nil
		}
		return time.Time{}, errors.New("empty value")
	}

	// Try absolute RFC 3339 first.
	if t, err := time.Parse(time.RFC3339, in); err == nil {
		return t.UTC(), nil
	}

	// Otherwise interpret as relative duration.
	dur, err := ParseHumanDuration(in)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().Add(-dur).UTC(), nil
}
