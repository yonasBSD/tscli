package file

import (
	"encoding/json"
	"github.com/tailscale/hujson"
)

func ValidatePolicy(b []byte) error {
	if json.Valid(b) {
		return nil
	}
	_, err := hujson.Parse(b)
	return err
}
