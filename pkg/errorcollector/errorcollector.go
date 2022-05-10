package errorcollector

import (
	"errors"
	"fmt"
	"strings"
)

type Fields map[string]error

func (f Fields) Add(field string, err error) { f[field] = err }

func (f Fields) Err() error {
	if len(f) > 0 {
		return errors.New(f.Error())
	}
	return nil
}

func (f Fields) Error() string {
	errstrings := make([]string, len(f), len(f))
	counter := 0
	for k, v := range f {
		errstrings[counter] = fmt.Sprintf("%s: %s", k, v.Error())
		counter++
	}
	return strings.Join(errstrings, "; ")
}
