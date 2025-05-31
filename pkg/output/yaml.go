package output

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v3"
)

type YAMLPrinter struct{}

func (YAMLPrinter) Print(b []byte) error {
	var obj any
	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}
	y, err := yaml.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(y)
	return err
}
