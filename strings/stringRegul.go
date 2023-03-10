package main

import (
	"fmt"
	"regexp"
)

func countRusWords(text string) int {
	var words bool
	var count int
	for _, val := range text {
		matched, _ := regexp.MatchString(`^[а-я]+$`, string(val))
		if !matched && string(val) != " " {
			words = false
			continue
		}
		if val == 32 && words == true {
			count++
		} else if val == 32 && !words {
			words = true
		}
	}
	return count
}

func main() {
	text := "Привет мир golong helloарас стекло пылеwсо dваа ванна клепа Привет"
	fmt.Println(countRusWords(text))
}
