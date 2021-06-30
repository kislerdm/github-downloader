package external

import "strings"

type Errors []error

func (e Errors) Error() string {
	o := []string{}
	for _, err := range e {
		o = append(o, err.Error())
	}
	return strings.Join(o, "\n")
}
