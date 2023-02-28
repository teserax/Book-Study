//Введите массив символов их 23 элементов.Определить, является ли он
//палиндромом(симметричным с точностью до пробелов)или нет.Например ,
//А РОЗА УПАЛА НА ЛАПУ АЗОРА - палиндром.
//(Предполагается, что все буквы строки - прописные.)

package main

import (
	"strings"
)

func palidrom(str string) bool {
	if len(str) == 0 {
		return false
	}
	str = strings.ReplaceAll(str, " ", "")//ВОПРОС должны мы уберать пробелы или нет???
	text := []rune(str)
	for i := 0; i < len(text); i++ {
		if text[i] != text[len(text)-1-i] { //если не совпадают то не палидром и выходим из цикла
			return false
		} else if len(text)-1-i < len(text)/2 { //если мы проверили больше половины длины слова значит все совподает ПАЛИДРОМ!
			return true
		}
	}
	return false
}
func main() {
	text := "А РОЗА УПАЛА НА ЛАПУ АЗОРА" //так  верно
	text = strings.ReplaceAll(text, " ", "")
	//text := []rune("А РОЗА УПАЛА НА ЛАПУ АЗОРА") //соответственно так нет верно???

}
