package validation

import (
	"bytes"
	"fmt"
)

type Result map[string]error

func (r Result) Error() string {
	var buf bytes.Buffer
	for field, err := range r {
		message := fmt.Sprintf("\"%s\": %v", field, err)
		buf.WriteString(message)
	}
	return buf.String()
}

func (r Result) Succeeded() bool {
	for _, errors := range r {
		if len(errors) > 0 {
			return false
		}
	}
	return true
}
