package file

import (
	"github.com/tailscale/hujson"
	"encoding/json"
)


func ValidatePolicy(b []byte) error {
	if json.Valid(b) {
		return nil
	}
	_, err := hujson.Parse(b)
	return err
}