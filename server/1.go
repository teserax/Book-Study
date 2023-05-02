// Основы сетевого программирования на Go - 12  глава
// с 580 по 604
// разобраться с кодом `Создание веб-сервера на Go`. Сроки - 5 дней.
// Прийдется чуть зацепить 11 главу -  профилирование
package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, i := range interfaces {
		fmt.Printf("Interface: %v\n", i.Name)
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}
		addresses, err := byName.Addrs()
		for k, v := range addresses {
			fmt.Printf("Interface Address #%v: %v\n", k, v.String())
		}
		fmt.Println()
	}
}