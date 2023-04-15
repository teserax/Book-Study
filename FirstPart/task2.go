/*Вопросы (только с условием ответ иcкать только в докумнетации пакета time):
- как к какой-то дате добавить отрезок времени. Например 1 час и 5 минут. Привести пример
- как зная строковое значение времени, например 5min - начать работать с ним как с интервалом времени (привести пример)
- как найти количество дней, лет, месяцев, часов между 2мя датами (привести пример)
- как посчитать и вывести на экран время сколько работала твоя программа (или функция).привести пример.
- как ставить паузу в работе программы (привести пример)*/

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	//- как к какой-то дате добавить отрезок времени.например 1 час и 5 минут. Привести пример
	//скажем у нас произвольно есть какая-то дата "Временной отрезок" у которого тип данных time.Time
	//у которого есть много удобных методов для работы с временным интервалом"датой"
	//один из методов Add() позволяет дабовлять кнашей дате временной отрезок
	//пример:
	var myDate time.Time
	myDate = time.Date(1980, 07, 23, 13, 15, 10, 0, time.UTC) //произвольные дата и время
	fmt.Printf(" Arbitrary date    %s \n", myDate.String())
	dateAfterChange := myDate.Add(time.Hour * 1).Add(time.Minute * 5) //добавляем один час и 5 минут
	fmt.Printf(" Date after change %s ", dateAfterChange.String())
	timeStr := "3h10m"                    //условное время для простоты в строковом представлении
	t, err := time.ParseDuration(timeStr) //парсим из строки в Duration тобишь(int64)
	if err != nil {
		fmt.Printf("An error occurred: %v.\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n Сonditional \" 3h10m \" time in minutes %.0f \n Сonditional time\" 3h10m \" in seconds %.0f \n ", t.Minutes(), t.Seconds())
	//благодоря методам .Minutes(), .Seconds() и тд выводим в интересующем (нужном нам представлении) данный участок времени
	workTime := time.Now()
	fmt.Println(workTime)                                                  //создаем переменную time.Time хранящую настоящее время благодаря методу Now()
	fmt.Printf("Time now is %d:%d \n", workTime.Hour(), workTime.Minute()) //так же блогодоря встроеным методам мы можем вывести
	// отдельно время, минуты,дни и тд
	workTime = workTime.Add(t) //используя метод Add можно добавить любой отрезок времени
	fmt.Printf(" Time now plus conditional time in minutes is %d:%d \n", workTime.Hour(), workTime.Minute())
	//выводим время с добавленным отрезком!!
	t1 := workTime.Sub(myDate).Minutes() / 60 //выводим время в часах
	fmt.Printf(" Elapsed time in Hours %.0f\n", t1)
	t1 = workTime.Sub(myDate).Hours() / 24 //выводим в днях
	fmt.Printf(" Elapsed time in Days  %.0f\n", t1)
	t1 = workTime.Sub(myDate).Hours() / 24 / 365 //выводим в годах
	fmt.Printf(" Elapsed time in Years %.0f\n", t1)
	end := time.Now().Sub(start) //время работы программ относительно прошедшему времени с начала старта программы
	fmt.Printf(" Time of work %v Seconds\n", end.Seconds())
	fmt.Println(" We fall asleep on 3 Seconds")
	time.Sleep(time.Second * 3) //задаем и останавливаем программу на заданное время
	end = time.Now().Sub(start)
	fmt.Printf(" Total running time of programs %v Seconds", end.Seconds()) //общее время всей программы
}
