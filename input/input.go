package input

import (
	"math"
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

// Stats struct for statistics on given input.
type Stats struct {
	ShortestWord                  string
	LongestWord                   string
	AverageWordLength             float64
	ListOfWordWithLengthAsAverage []string
}

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

// WholeStory returns text that is composed from all words separated by spaces.
func (s Split) WholeStory() string {
	if len(s) == 0 {
		return ""
	}

	var res []string
	for _, ss := range s {
		res = append(res, ss.Word)
	}
	return strings.Join(res, " ")
}

// StoryStats returns statistics on given input.
func (s Split) StoryStats() Stats {
	var shortestWord, longestWord string
	var lengthShortest, lengthLongest, totalLength, averageWordLengthRounded int
	var averageWordLength float64
	var listOfWordWithLengthAsAverage []string
	length := len(s)

	if length == 0 {
		return Stats{}
	}

	lengthShortest = len(s[0].Word)
	for _, ss := range s {
		if len(ss.Word) > lengthLongest {
			longestWord, lengthLongest = ss.Word, len(ss.Word)
		}
		if len(ss.Word) <= lengthShortest {
			shortestWord, lengthShortest = ss.Word, len(ss.Word)
		}
		totalLength += len(ss.Word)
	}
	averageWordLength = float64(totalLength) / float64(length) // 0.5
	// rounded up and down:
	// if the number is followed by 5, 6, 7, 8, or 9, round the number up;
	// if the number is followed by 0, 1, 2, 3, or 4, round the number down.
	averageWordLengthRounded = int(math.Floor(averageWordLength + 0.5))

	for _, ss := range s {
		if len(ss.Word) == averageWordLengthRounded {
			listOfWordWithLengthAsAverage = append(listOfWordWithLengthAsAverage, ss.Word)
		}
	}

	return Stats{
		shortestWord,
		longestWord,
		averageWordLength,
		listOfWordWithLengthAsAverage,
	}
}
