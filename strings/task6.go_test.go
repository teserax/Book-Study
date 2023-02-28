package main

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	var tests = []struct {
		text   string
		string string
	}{
		{"a=+1", "a++"},
		{"a:=1", "a:=1"},
		{"b=-1", "b--"},
		{"a=+1;b:=-1", "a++;b:=-1"},
	}
	for _, test := range tests {
		if got := iterator(test.text); got != test.string {
			fmt.Printf("%v %v", got, test.string)
			t.Errorf(`Error!  want: %v  got: %v`, test.string, got)
		}
	}
}
