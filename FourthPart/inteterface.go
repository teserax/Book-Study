//Из главы (8. Как объяснить UNIX-системе, что она должна делать)
//Задача для этого: Напишите программу Go, которая принимает три аргумента:
//"имя текстового файла" и "две строки". Затем эта утилита должна заменить каждое вхождение
//первой строки в файле второй строкой. В целях безопасности окончательный
//результат выведите на экран, чтобы исходный текстовый файл остался без из-
//менений. (2 дня)

package main

import "fmt"

type makeVoice interface {
	Voice()
}

type human struct{}   //разные сущности
type animal struct{}  //разные сущности
type vehicle struct{} //разные сущности

func (v vehicle) Voice() {
	fmt.Println("bbrrr,bbrrr")
}
func (a animal) Voice() {
	fmt.Println("animal roars ra ra ra")
}
func (h human) Voice() {
	fmt.Println("Human say ha ha ha")
}
func voice(m makeVoice) { //блогодоря интерфейсу мы можем предать любую сущность у которой есть метод Voice
	m.Voice()
}

func main() {
	andrei := human{}
	cat := animal{}
	GT90 := vehicle{}
	voice(andrei)
	voice(cat)
	voice(GT90)
	//можем добавить или удалить любую сущность
}
