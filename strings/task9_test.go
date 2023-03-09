package main

import (
	"fmt"
	"testing"
)

func TestFindRussianCapitals(t *testing.T) {
	var tests = []struct {
		text string
		want int
	}{{"Привет мир golong helloарас стекло пылеwсо", 2},
		{"Привет  пылеwсо", 0},
		{"Привет мир golong helloарас стеwкло пылеwсо тwеwсwт", 1},
	}
	for _, test := range tests {
		if got := findRussianCapitals(test.text); got != test.want {
			fmt.Errorf(`Error!  want: %v  got: %v`, test.text, got)
		}
	}
}
