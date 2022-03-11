package input

import (
	"regexp"
	"strconv"
	"strings"
)

// Node struct for number-word node from given input.
type Node struct {
	Number uint
	Word   string
}

// Split slice of number-word nodes.
type Split []Node

var inputRegexp = regexp.MustCompile(`^([0-9]+-[^-]+)(-[0-9]+-[^-]+)*$`)

// TestValidity tells is given string valid or not.
func TestValidity(s string) bool {
	return inputRegexp.MatchString(s)
}

// FromString takes given input and parses it to slice of number-word nodes.
func FromString(s string) Split {
	data := strings.Split(s, "-")
	res := make(Split, 0)

	if data[0] == s {
		return res
	}

	for i, ss := range data {
		if i%2 == 0 {
			ssInt, _ := strconv.Atoi(ss)
			res = append(res, Node{uint(ssInt), data[i+1]})
		}
	}

	return res
}

// AverageNumber returns the average number from all the numbers.
func (s Split) AverageNumber() float64 {
	length := len(s)
	if length == 0 {
		return 0
	}

	var total uint
	for _, ss := range s {
		total += ss.Number
	}
	return float64(total) / float64(length)
}
