package file

import (
	"io"
	"os"
	"strings"
)

func ReadInput(path, inline string) ([]byte, error) {
	if inline != "" {
		return []byte(inline), nil
	}
	if path == "-" {
		return io.ReadAll(os.Stdin)
	}
	if strings.HasPrefix(path, "file://") {
		path = strings.TrimPrefix(path, "file://")
	}
	return os.ReadFile(path)
}
