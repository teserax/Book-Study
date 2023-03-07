package main

import (
	"fmt"
	"testing"
)

func TestNumberOfSimbols(t *testing.T) {
	var tests = []struct {
		text       string
		lengthWord int
		want       int
	}{{"hello go hello", 0, 0},
		{"hello o hello o", 1, 5},
		{"hello go hello", 5, 2},
		{"hello go hello go", 2, 5},
		{"hello go hello go firsts hello", 5, 8},
	}
	for _, test := range tests {
		if got := numberOfSimbols(test.text, test.lengthWord); got != test.want {
			fmt.Errorf(`Error!  want: %v  got: %v`, test.want, got)
		}
	}
}
