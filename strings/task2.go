//Есть некий текст. нужно посчитать количество слов в этом тексте.
//Если известно что слова разделяются одним пробелом.

package main

import "fmt"

func main() {
	text := "Ut enim ad minim veniam nam libero tempore cum soluta nobis est eligendi optio cumque nihil impedit quo minus id"
	count := 1
	for _, val := range text {
		if string(val) == " " {
			count++
		}
	}
	fmt.Println(count)
}
