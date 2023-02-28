//Есть некий текст. Предложения заканчиваются на знаках ".", "!", "?".
//Нужно посчитать количество предложений и вывести на экран их в обратном порядке.

package main

import "fmt"

func main() {
	text := "первый! второй stop? forever iang! пятый, No-6! rdggrs fsrsfs final."
	for i := 0; i < len(text); i++ {
		switch {
		case string(text[i]) == "!":
			text = text[:i] + text[i+1:]
			i--
		case string(text[i]) == "?":
			text = text[:i] + text[i+1:]
			i--
		case string(text[i]) == ",":
			text = text[:i] + text[i+1:]
			i--
		case string(text[i]) == ".":
			text = text[:i] + text[i+1:]
			i--
		}

	}
	fmt.Println(text)
}
