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
		{"first second", "second first"},
	}

	for _, test := range test {
		if got := reversText(test.testText); got != test.returnText {
			fmt.Printf("%v %v", got, test.returnText)
			t.Errorf(`Error!  want: %v  got: %v`, test.returnText, got)
		}

	}

}
