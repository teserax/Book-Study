package main

import "testing"

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
		if got := countWordsInText(test.text); got  test.want {
			t.Error("Error! countWordsInText() ", test.text, got)
		}
	}

}
