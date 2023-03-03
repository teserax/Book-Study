package main

import (
	"fmt"
	"testing"
)

func TestReversText(t *testing.T) {
	var test = []struct {
		testText   string
		returnText string
	}{
		{"first", "first"},
		{"first second", "second first"},
		{"first second first", "second first first"},
		{"first second first second", "second first second first"},
		{"первый второй первый второй", "второй первый второй первый"},
		{"первый второй первый", "второй первый первый"},
	}

	for _, test := range test {
		if got := reversText(test.testText); got != test.returnText {
			fmt.Printf("%v %v", got, test.returnText)
			t.Errorf(`Error!  want: %v  got: %v`, test.returnText, got)
		}

	}

}
