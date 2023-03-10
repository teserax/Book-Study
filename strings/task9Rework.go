//Решить в одном обходе строки,суть решение - идти по буквам.проверять если русская, учитывать пробелы как начало и конец слова

package main

import (
	"fmt"
	"unicode"
)

func countRussionWords(text string) int {
	var count, count1, count2 int //задаем 3 счетчика
	for _, val := range text {
		count1++                               //первый считает общее количество символов до пробела
		s := unicode.Is(unicode.Cyrillic, val) //проверка на то что символ является буквой кириллицы
		if s == true && unicode.IsLower(val) { //проверяем что буква в нижнем регистре
			count2++ //второй счетчик подсчитывает количество вхождений искомых символов
		}
		//согласно коду ASCII пробелу соответствует порядковый номер 32
		if val == 32 { //если пробел то
			count1--              //от основного счетчика символов минус 1 ..пробел..
			if count1 == count2 { //если количество искомых символов совпадает с пройденными до пробела
				count++ //значит группа символ от начала и до конца полностью отвечает заданным параметрам
				//добавляем в основной счетчик плюс 1
			}
			count1 = 0 //обнуляем счетчики
			count2 = 0 //обнуляем счетчики
		}

	}
	return count //возвращаем количество найденных слов

}
func main() {
	text := "Привет мир golong helloарас стекло пылеwсо dваа ванна клепа Привет"
	fmt.Println(countRussionWords(text))
}