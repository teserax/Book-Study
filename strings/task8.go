//В заданной строке расположить в обратном порядке все слова.
//Разделителями слов считаются пробелы.
package main

import (
	"fmt"
	"strings"
)

func reversText(text string) string {
	if len(text) == 0 {
		return text
	}
	splitText := strings.Split(text, " ")
	if len(splitText) == 1 {
		return text
	}
	r := []string{}
	i := len(splitText)
	for i > 0 {
		i--
		r = append(r, splitText[i])
	}
	result := strings.Join(r, " ")
	return result
}
func main() {
	//text := "first second third fourth"
	text := "Первый второй третий четвертый"
	fmt.Println(reversText(text))

}
