/*
Ведите массив символов из 17 элементов, содержащих два или более слов,
разделенных пробелами. Поменять местами все четные и не четные слова в массиве
*/

package main

import (
	"fmt"
	"strings"
)

func reversText(text string) string {
	splitText := strings.Split(text, " ")
	var result string
	for i := 0; i < len(splitText); i++ {
		if i+1 == len(splitText) {
			result = " " + result + splitText[i]
		} else {
			result = result + splitText[i+1] + " " + splitText[i] + " "
			i++
		}

	}
	result = strings.TrimSpace(result)
	return result
}
func main() {

	//text := "Temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet, unde omnis iste natus error"
	//	text := "first second first second first second first second"
	text := "первый второй первый второй"
	fmt.Println(len(text), len(reversText(text)))
	fmt.Println(text, reversText(text))
}
