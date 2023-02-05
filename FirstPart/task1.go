/*Напишите программу кторая вычисляет сумму всех аргументов командной строки,
которые являются действительными числами*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 { //если длина переданных аргументов ровна "1" это означает что нам ничего не передали
		//  так как первое значение  в этом срезе - это путь к программе
		fmt.Println("please give me one argument!")
		return
	}
	total := 0
	for i := 1; i < len(os.Args); i++ { //если передали.... то используя цикл for мы передаем полученные аргумент в переменную number
		//и спользуя метотд пакета-strconv.Atoi конвертируем полученные данные в число так как  мы получаем их виде string
		number, err := strconv.Atoi(os.Args[i])
		if err != nil { //проверяем на корректность данных
			fmt.Println(err)
			return //если получим ошибку выводим в консоль и выходи из программы
		}
		total += number //суммируем
	}

	fmt.Println(total) //выводим полученный результат в консоль
}
