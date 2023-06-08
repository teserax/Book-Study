// Есть некий текстовый файл и в нем много текста.
// Нужно сравнить скорость работы с текстом
// через работу с текстом кодом и при помощи рег выражений.
// Например можно сделать:
// - подсчет гласных букв
// - подсчет количества email адресов в тексте
// надо написать 2 функции.
// Одна задачу делает без рег выражений,
// вторая С рег. выражением.
// И сравнить бенчмарком скорость
// работы обоих функций и сделать выводы
package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

const strVowels = "aeiouy" //гласные

func SearchVowels(str string) int { //поиск гласных
	if str == "" {
		return 0
	}
	count := 0
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(strVowels); j++ {
			if str[i] == strVowels[j] {
				count++
			}
		}
	}

	return count
}
func SearchVowelsInSlice(str string) int { //поиск гласных хранящихся в слайсе
	strVowels := []string{"a", "e", "i", "o", "u", "y"}
	count := 0
	for _, val := range str {
		for _, word := range strVowels {
			if string(val) == word {
				count++
			}
		}
	}
	fmt.Println(count)
	return count
}
func SearchVowelsRegular(str string) int { //поиск гласных через регулярные выражения
	if str == "" {
		return 0
	}
	count := 0
	for _, word := range strVowels {
		re, err := regexp.Compile(string(word))
		if err != nil {
			log.Fatal(err)
		}
		res := re.FindAllString(str, -1)
		count += len(res)
	}
	return count
}
func SearchVowelsInMap(m map[string]int) int { //поиск гласных в мапе
	if len(m) == 0 {
		return 0
	}
	count := 0
	for word, num := range m {
		for _, vowels := range strVowels {
			if word == string(vowels) {
				count += num
			}
		}

	}

	return count
}

func ReadFile(nameFile string) string { // функция для чтения файла
	bytes, err := os.ReadFile(nameFile)
	if err != nil {
		log.Fatal(err)
	}
	fileText := string(bytes[:])

	return fileText
}
func stringToMap(str string) map[string]int { //конвертация строки в мапу
	m := map[string]int{}
	for _, word := range str {
		if _, ok := m[string(word)]; ok {
			m[string(word)]++
		} else {
			m[string(word)]++
		}
	}
	return m
}
func main() {

	ReadFile("text.txt")
	fmt.Println(SearchVowels(ReadFile("text.txt")))                   //count 2262735
	fmt.Println(SearchVowelsRegular(ReadFile("text.txt")))            //count 2262735
	fmt.Println(SearchVowelsInMap(stringToMap(ReadFile("text.txt")))) //count 2262735
	fmt.Println(SearchVowelsInSlice(ReadFile("text.txt")))            //count 2262735
}
