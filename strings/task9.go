//В заданной строке посчитать количество слов,
//содержащих только строчные русские буквы.
//Разделителями слов считаются пробелы

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func findRussianCapitals(text string) string {
	sText := strings.Split(text, " ")
	result := []string{}
	for _, val := range sText {
		matched, _ := regexp.MatchString(`^[а-я]+$`, val) //тут разобрался"^"сначала строки ,"+"все заданные символы и "$" доконца строки...тут слово
		if matched == true {
			result = append(result, val)
		}
	}
	return strings.Join(result, " ")
}
func main() {
	text := "" //запутался как исключить слова в которых есть англиские буквы как последннее??"пылесоw
	fmt.Println(findRussianCapitals(text))
}
