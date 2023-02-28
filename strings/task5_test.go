package main

import (
	"testing"
)

func TestPalidrom(t *testing.T) {
	var tests = []struct {
		text string
		pali bool
	}{
		{"топот", true},
		{"тоerrgeот", false},
		{"А РОЗА УПАЛА НА ЛАПУ АЗОРА", true},
	}
	for _, test := range tests {
		if got := palidrom(test.text); got != test.pali {
			t.Errorf(`Error!  want: %v  got: %v`, test.pali, got)
		}
	}

}

