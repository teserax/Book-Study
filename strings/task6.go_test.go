package main

import "testing"

func TestIterator(t *testing.T) {
	var tests = []struct {
		text   string
		string string
	}{
		{"a=+1", "a++"},
		{"a:=1", "a:=1"},
		{"b=-1,a=+1", "b--,a++"},
	}
	for _, test := range tests {
		if got := palidrom(test.text); got == test.string {
			t.Errorf(`Error!  want: %v  got: %v`, test.string, got)
		}
	}

}
