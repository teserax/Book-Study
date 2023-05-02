package main

import (
	"fmt"
	"testing"
)

func TestSearchVowels(t *testing.T) {
	var tests = []struct {
		testText string
		count    int
	}{{"", 0},
		{"first", 1},
		{"first second third fourth", 6},
		{"первый второй ", 3},
		{"first secondaeiouy", 9},
	}
	for _, test := range tests {
		if got := SearchVowels(test.testText); got != test.count {
			fmt.Errorf(`Error!  want: %v  got: %v`, test.count, got)
		}
	}
}
func TestSearchVowelsRegular(t *testing.T) {
	var tests = []struct {
		testText string
		count    int
	}{{"", 0},
		{"first", 1},
		{"first second third fourth", 6},
		{"первый второй ", 3},
		{"first secondaeiouy", 9},
	}
	for _, test := range tests {
		if got := SearchVowelsRegular(test.testText); got != test.count {
			fmt.Errorf(`Error!  want: %v  got: %v`, test.count, got)
		}
	}
}
