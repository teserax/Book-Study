package main

import (
	"fmt"
	"testing"
)

func TestFindRussianCapitals(t *testing.T) {
	var tests = []struct {
		text string
		want string
	}{{"Привет мир golong helloарас стекло пылеwсо", "мир стекло"},
		{"Привет  пылеwсо", ""},
		{"Привет мир golong helloарас стеwкло пылеwсо тwеwсwт", "мир"},
	}
	for _, test := range tests {
		if got := findRussianCapitals(test.text); got != test.want {
			fmt.Errorf(`Error!  want: %v  got: %v`, test.text, got)
		}
	}
}
