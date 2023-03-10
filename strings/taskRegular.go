package main

import (
	"fmt"
	"unicode"
)

func main() {
	text := "Привет мир golong helloарас Стекло пылеwсо dваа ванна клепа Привет"
	var words bool
	var count int
	for _, val := range text {

		if (unicode.Is(unicode.Cyrillic, val) && unicode.IsUpper(val)) || !unicode.Is(unicode.Cyrillic, val) && string(val) != " " {
			words = false

			continue
		}
		if val == 32 && words == true {
			count++
		} else if val == 32 && !words {
			words = true
		}
	}
	fmt.Println(count)
}
