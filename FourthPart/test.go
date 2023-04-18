//Из главы (8. Как объяснить UNIX-системе, что она должна делать)
//Задача для этого: Напишите программу Go, которая принимает три аргумента:
//"имя текстового файла" и "две строки". Затем эта утилита должна заменить каждое вхождение
//первой строки в файле второй строкой. В целях безопасности окончательный
//результат выведите на экран, чтобы исходный текстовый файл остался без из-
//менений. (2 дня)

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := flag.String("filename", "text", "The name")                    //флаг на имя файла по умолчанию
	firstStr := flag.String("first", "Vova", "The First string")               //флаг на изначальное слова для замены
	replacString := flag.String("replace", "Valera", "The Replacement string") // флаг на слово для замены
	flag.Parse()
	*filename += ".txt"                 //добавляю разрешение файла
	file, err := os.ReadFile(*filename) //читаем
	if err != nil {
		fmt.Printf("Open file error %s", err) //ошибка
	}
	fmt.Println(strings.ReplaceAll(string(file), *firstStr, *replacString)) //меняем и выводим
	fmt.Printf("Text on file %s\n", file)                                   // основной файл неизменен
}
