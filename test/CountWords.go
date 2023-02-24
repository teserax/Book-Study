// Напишите функцию для подсчета частоты слов в текстовой строке и возврата карты
// слов с их количеством. Функция должна преобразовывать текст в нижний регистр, и
// знаки препинания следует отделять от слов. Пакет strings содержит несколько справок:
// Полные функции для этой задачи, включая Fields, ToLower и Trim.
// Используйте свою функцию, чтобы подсчитать частоту слов в следующем отрывке, а затем
// отображать количество слов, которые встречаются более одного раза.

package main

import (
	"fmt"
	"strings"
)

func сountWordsInText(text string) map[string]int {
	textSymbols := `!,$%^&*()/[{}.`
	if len(text) == 0 {
		return map[string]int{}
	}
	for _, val := range text { //проверка и очистка текста на спец-символы
		for i := 0; i < len(textSymbols); i++ {
			if val == int32(textSymbols[i]) { //если найден удаляем и приводим строку к нижнему регистру согласно условиям
				text = strings.ReplaceAll(strings.ToLower(text), string(textSymbols[i]), "")
			}
		}
	}
	words := map[string]int{}
	textSplit := strings.Split(text, " ")
	for _, val := range textSplit {
		if _, exist := words[string(val)]; exist {
			words[string(val)]++
		} else {
			words[string(val)] = 1
		}
	}
	return words
}
func main() {
	text := "It has survived not [[only! five centuries, bu[t also the leap into electronic[[ typesetting, five remaining[[ essentially unchanged."
	var textRu string
	textRu = "Мы приложили много усилий, чтобы работать на нашей бирже копирайтинга было максимально легко и удобно"
	fmt.Println(сountWordsInText(text))
	fmt.Println(сountWordsInText(textRu))

}