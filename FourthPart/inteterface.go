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
	"strings"
)

type Value interface {
	String() string
	Set(string) error
}

type NamesFlag struct {
	Names []string
}

func (s *NamesFlag) GetNames() []string {
	return s.Names
}
func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}
func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("Cannot use names flag more than once!")
	}
	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}
func main() {
	var manyNames NamesFlag
	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "Mihalis", "The name")
	flag.Var(&manyNames, "names", "Comma-separated list")
	flag.Parse()
	fmt.Println("-k:", *minusK)
	fmt.Println("-o:", *minusO)
	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}
	fmt.Println("Remaining command line arguments:")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}
