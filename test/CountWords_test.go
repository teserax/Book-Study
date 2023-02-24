package main

import (
	"reflect"
	"testing"
)

func TestCountWordsInText(t *testing.T) {
	var tests = []struct {
		text string
		want map[string]int
	}{
		{"", map[string]int{}},
		{"hi world", map[string]int{"hi": 1, "world": 1}},
		{"hi ,hi world!", map[string]int{"hi": 2, "world": 1}},
	}
	for _, test := range tests {
		if got := сountWordsInText(test.text); reflect.DeepEqual(got, test.want) != true {
			t.Errorf(`Error! countWordsInText() `, test.want, got)
		}
	}

}
