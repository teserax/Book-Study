//Есть некий текст в котором представлены какие то числа. Пример текста:
//Текст содержит числа. Нужно вытянуть все числа, отсортировать по возрастанию и вывести на экран результат.

/*Я родился 23 ноября 1981 года. Я живу по адресу Болгарская 43, квартира 34. У меня 3 детей.*/

package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	text := "Я родился 23 ноября 1981 года. Я живу по адресу Болгарская 43, квартира 34. У меня 3 детей."
	re := regexp.MustCompile("[0-9]+")
	s := re.FindAllString(text, -1)
	numbers := []int{}
	for i := 0; i < len(s); i++ {
		n, err := strconv.Atoi(s[i])
		if err != nil {
			fmt.Println(err)
		} else {
			numbers = append(numbers, n)
		}
	}
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i] < numbers[j] {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}
	fmt.Println(numbers)
}
