//Есть некий текст. Предложения заканчиваются на знаках ".", "!", "?".
//Нужно посчитать количество предложений и вывести на экран их в обратном порядке.

package main

import "fmt"

func main() {
	text := "первый!второй stop?forever iang!пятый.No-6!rdggrs fsrsfs final."
	str := []string{}
	for i := 0; i < len(text); i++ {
		switch {
		case string(text[i]) == "!":
			str = append(str, text[:i+1])
			text = text[i+1:]
			i = 0
		case string(text[i]) == "?":
			str = append(str, text[:i+1])
			text = text[i+1:]
			i = 0
		case string(text[i]) == ".":
			str = append(str, text[:i+1])
			text = text[i+1:]
			i = 0
		}

	}
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Println(str[i])
	}
}
