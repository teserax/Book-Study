package main

import (
	"fmt"
	"testing"
)

func TestCountRusWords(t *testing.T) {
	var tests = []struct {
		text string
		want int
	}{
		{"", 0},
		{"Good", 0},
		{"Good день", 1},
		{"Хороший День", 0},
		{"Привет мир golong helloарас стекло пылеwсо dваа ванна клепа Привет", 4},
	}
	for _, test := range tests {
		if got := countRusWords(test.text); got != test.want {
			fmt.Errorf("want %d ,got %d", test.want, got)
		}
	}
}
