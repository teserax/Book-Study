//Даны 5 массивов символов, представляющих фрагмент текста программы.
//Известно, что данный фрагмент оперирует только с целочисленными переменными
//В каждой строке-одна команда
//Преобразовать данный тест заменив каждую строку вида
//          /переменная=переменная "+1" на строку "++"/
// А каждую строку вида
//          /переменная=переменная "-1" на строку переменная "--"/

package main

import "fmt"

func iterator(text string) string {
	if len(text) <= 1 {
		return text
	}
	for i := 0; i < len(text); i++ {
		if string(text[i]) == ":" {
			return text
		}
		if string(text[i]) == "=" && string(text[i+1]) == "+" && string(text[i+2]) == "1" {
			text = text[:i] + "++" + text[i+3:]
		}
		if string(text[i]) == "=" && string(text[i+1]) == "-" && string(text[i+2]) == "1" {
			text = text[:i] + "--" + text[i+3:]
		}

	}
	return text
}
func main() {
	text := "a=+1"
	fmt.Println(iterator(text))
}
