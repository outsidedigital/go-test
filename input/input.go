package input

import (
	"regexp"
)

var inputRegexp = regexp.MustCompile(`^([0-9]+-[^-]+)(-[0-9]+-[^-]+)*$`)

// TestValidity tells is given string valid or not.
func TestValidity(s string) bool {
	return inputRegexp.MatchString(s)
}
