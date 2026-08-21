package jsons

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Unmarshal(content []byte, object any) error {
	first := bytes.Index(content, []byte{'{'})
	last := bytes.LastIndex(content, []byte{'}'})
	if first == -1 || last == -1 {
		return fmt.Errorf("object not found")
	}

	err := json.Unmarshal(content[first:last+1], object)
	if err != nil {
		return err
	}

	return nil
}
