package main

import "testing"

func BenchmarkSearchVowelsCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchVowels(ReadFile("text.txt"))
	}
}
func BenchmarkSearchVowelsRegular(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchVowelsRegular(ReadFile("text.txt"))
	}
}
