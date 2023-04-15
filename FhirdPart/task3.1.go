// -пусть есть некий слайс структур, в котором описаны
// какие то события в биографии человека.
// Каждая структура имеет 2 поля: дата и событие.
// Притом даты указаны в  одном из 4-х форматов:
// .......... (пример: 2023-01-22) yyyy-mm-dd
// .......... (пример 2022-01-22 10:00:00,
// .......... (пример: 01/12/2022),dd/mm/yyyy
// .......... (пример: Mon, 02 sep 2022 15:04:05).
// Нужно написать функцию которая бы возвращала события в отсортированом виде.
// Функция  также должна принимать в качестве флага вид сортировки:
// от наибольщего к меньшему или наоборот.
// Также написать тест к функции
package main

import (
	"flag"
	"fmt"
	"time"
)

type event struct {
	date  string
	event string
}
type eventWithDate struct {
	date  time.Time
	event string
}

func main() {
	myEvents := []event{
		{
			date:  "2024-02-07",
			event: "birthday",
		},
		{
			date:  "01-12-2023",
			event: "enlightenment day",
		},
		{
			date:  "02 sep 2025",
			event: "significant day",
		},
		{
			date:  "07/07/1981",
			event: "significant day",
		},
	}
	layouts := []string{"02/01/2006", "02 Jan 2006", "02/01/06", "2006-01-02", "02-01-2006"}
	// единый формат дат в слайсе. time.Time - это тип времени, которы ей представляет собой структуры с своими методами (см. докумнетацию покета time).
	sortEvent := make([]eventWithDate, len(myEvents))

	for i, event := range myEvents {
		var err error
		for _, layout := range layouts {
			var t time.Time
			t, err = time.Parse(layout, event.date)
			if err == nil {
				sortEvent[i].date = t
				sortEvent[i].event = event.event
				break
			}
		}
		if err != nil {
			fmt.Printf("for date %q, supported formats: %#v, error: %v\n", event.date, layouts, err)
		}
	}
	sort := flag.Bool("sort", false, "sort date ")
	flag.Parse()
	sortedEvents := sortDates(sortEvent, *sort)
	for _, event := range sortedEvents {
		fmt.Printf("%+v\n", event.date.String())
	}
}
func sortDates(myEvents []eventWithDate, sort bool) []eventWithDate {
	for i := 0; i < len(myEvents); i++ {
		for j := i + 1; j < len(myEvents); j++ {
			if sort && myEvents[j].date.Sub(myEvents[i].date) > time.Since(myEvents[i].date) {
				temp := myEvents[i].date
				myEvents[i].date = myEvents[j].date
				myEvents[j].date = temp
			} else if !sort && time.Since(myEvents[j].date) < time.Since(myEvents[i].date) {
				temp := myEvents[i].date
				myEvents[i].date = myEvents[j].date
				myEvents[j].date = temp
			}
		}
	}

	return myEvents
}
