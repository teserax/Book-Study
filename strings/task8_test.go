package main

import (
	"fmt"
	"testing"
)

func TestReversText(t *testing.T) {
	var tests = []struct {
		testText   string
		returnText string
	}{{"", ""},
		{"first", "first"},
		{"first second third fourth", "fourth third second first"},
		{"первый второй третий четвертый", "четвертый третий второй первый"},
		{"first second", "second first"},
	}
	for _, test := range tests {
		if got := reversText(test.testText); got != test.returnText {
			fmt.Errorf(`Error!  want: %v  got: %v`, test.returnText, got)
		}
	}
}
