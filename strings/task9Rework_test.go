package main

import (
	"fmt"
	"testing"
)

func TestCountRussionWords(t *testing.T) {
	var tests = []struct {
		text string
		want int
	}{{"", 0},
		{"Hello world", 0},
		{"Привет мир", 1},
		{"Привет мир world", 1},
		{"Привет миiр world", 0},
		{"привет мир, голанга", 3},
	}
	for _, test := range tests {
		if got := countRussionWords(test.text); got != test.want {
			fmt.Errorf(`Error!  want: %v  got: %v`, test.want, got)
		}
	}
}
