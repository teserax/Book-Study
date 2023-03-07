// Дан массив символов,содержащих текст.
// Расстояние между двумя словами равной длины-это количество символов,
// которыми они разделяются.
// В заданном массиве найти пару слов заданной длины с максимальным расстоянием.
package main

import (
	"fmt"
	"strings"
)

func numberOfSimbols(text string, length int) int {
	numSimbol := 0

	sText := strings.Split(text, " ")
	max := 0
	t := false
	for _, word := range sText {
		if length == len(word) {
			t = true
			//words = append(words, numSimbol)
			if numSimbol > max {
				max = numSimbol
			}
			numSimbol = 0
		} else if length != len(word) && t == true {
			numSimbol = numSimbol + len(word)
		}

	}
	return max
}

func main() {
	text := "hello wold sa s hello wold sa s fase hello wold wewr fsdsfsf sa s fase hello"
	fmt.Println("enter word length")
	var length int
	fmt.Scan(&length)
	fmt.Println(numberOfSimbols(text, length))
}
