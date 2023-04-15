package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Привет мир golong helloарас стекло пылеwсо dваа ванна клепа Привет"
	re := regexp.MustCompile("[а-я]+.")
	s := re.FindAllString(text, -1)
	fmt.Println(s)
}
