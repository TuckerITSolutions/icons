package utils

import "strings"

type MultiError []error

func (me MultiError) Error() string {
	if len(me) == 0 {
		return ""
	}
	// Join all individual error messages with a newline for readability
	s := make([]string, len(me))
	for i, err := range me {
		s[i] = err.Error()
	}
	return "Multiple errors occurred:\n" + strings.Join(s, "\n")
}
