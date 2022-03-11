package main

import (
	"reflect"
	"testing"

	"go-test/input"
)

func Test_testValidity(t *testing.T) {
	test := func(input string, want bool) func(*testing.T) {
		return func(t *testing.T) {
			got := testValidity(input)
			if got != want {
				t.Fatal("failed")
			}
		}
	}

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "example",
			input: "23-ab-48-caba-56-haha",
			want:  true,
		},
		{
			name:  "hex numbers",
			input: "0x25-",
			want:  false,
		},
		{
			name:  "empty",
			input: "",
			want:  false,
		},
		{
			name:  "empty string part",
			input: "10--10",
			want:  false,
		},
		{
			name:  "no dash",
			input: "23ab48caba56haha",
			want:  false,
		},
		{
			name:  "incomplete sequence",
			input: "23-ab-48-caba-56",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.input, tt.want))
	}
}

func Test_averageNumber(t *testing.T) {
	test := func(input string, want float64) func(*testing.T) {
		return func(t *testing.T) {
			got := averageNumber(input)
			if got != want {
				t.Fatal("failed")
			}
		}
	}

	tests := []struct {
		name  string
		input string
		want  float64
	}{
		{
			name:  "one number",
			input: "111-a",
			want:  111,
		},
		{
			name:  "multiple numbers",
			input: "111-aaa-222-bbb",
			want:  166.5,
		},
		{
			name:  "empty",
			input: "",
			want:  0,
		},
		{
			name:  "invalid",
			input: "23ab48caba56haha",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.input, tt.want))
	}
}

func Test_wholeStory(t *testing.T) {
	test := func(input string, want string) func(*testing.T) {
		return func(t *testing.T) {
			got := wholeStory(input)
			if got != want {
				t.Fatal("failed")
			}
		}
	}

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "one word",
			input: "111-a",
			want:  "a",
		},
		{
			name:  "multiple words",
			input: "111-aaa-222-bbb",
			want:  "aaa bbb",
		},
		{
			name:  "empty",
			input: "",
			want:  "",
		},
		{
			name:  "invalid",
			input: "23ab48caba56haha",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.input, tt.want))
	}
}

func Test_storyStats(t *testing.T) {
	test := func(input string, want input.Stats) func(*testing.T) {
		return func(t *testing.T) {
			got := storyStats(input)
			if !reflect.DeepEqual(got, want) {
				t.Fatal("failed")
			}
		}
	}

	tests := []struct {
		name  string
		input string
		want  input.Stats
	}{
		{
			name:  "one word",
			input: "111-a",
			want: input.Stats{
				ShortestWord:                  "a",
				LongestWord:                   "a",
				AverageWordLength:             1,
				ListOfWordWithLengthAsAverage: []string{"a"},
			},
		},
		{
			name:  "multiple words",
			input: "111-aaa-222-bb",
			want: input.Stats{
				ShortestWord:                  "bb",
				LongestWord:                   "aaa",
				AverageWordLength:             2.5,
				ListOfWordWithLengthAsAverage: []string{"aaa"},
			},
		},
		{
			name:  "empty",
			input: "",
			want:  input.Stats{},
		},
		{
			name:  "invalid",
			input: "23ab48caba56haha",
			want:  input.Stats{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.input, tt.want))
	}
}
