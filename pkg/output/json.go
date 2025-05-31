// pkg/output/json.go
package output

import (
	"encoding/json"
	"fmt"
	"os"
)

type JSONPrinter struct{}

func (JSONPrinter) Print(b []byte) error {
	var obj any
	if err := json.Unmarshal(b, &obj); err != nil {
		_, err = os.Stdout.Write(b)
		return err
	}

	out, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return err
	}

	_, err = fmt.Println(string(out))
	return err
}
