package main

import "go-test/input"

// testValidity tells is given string valid or not.
// Difficulty of the function: O(n)
// Time to implement the function: 20m
func testValidity(s string) bool {
	return input.TestValidity(s)
}

// averageNumber tells is given string valid or not.
// Difficulty of the function: O(n)
// Time to implement the function: 20m
func averageNumber(s string) float64 {
	i := input.FromString(s)
	return i.AverageNumber()
}

// wholeStory returns the average number from all the numbers.
// Difficulty of the function: O(n)
// Time to implement the function: 10m
func wholeStory(s string) string {
	i := input.FromString(s)
	return i.WholeStory()
}

// storyStats returns statistics on given input.
// Difficulty of the function: O(n)
// Time to implement the function: 10m
func storyStats(s string) input.Stats {
	return input.Stats{}
}
