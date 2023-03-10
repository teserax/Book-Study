package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "aaa=aaa+1"
	parts := strings.Split(str, "=")
	var result string
	// всегда надо проверять если есть элементы которые ты будешь по индексу доставать,
	// чтоб не попадать в панику при не правильной строке
	// проверки в голанге занимают ощутимую часть кода, так как себя обезопасить надо со всех сторон.
	if len(parts) > 1 {
		if partsPlus := strings.Split(parts[1], "+"); len(partsPlus) > 1 && partsPlus[0] == parts[0] {
			result = parts[0] + "++"
		} else {
			fmt.Printf("Incorrect string format (not contains +): %q", str)
			return
		}

	} else {
		fmt.Printf("Incorrect string format (not contains =): %q", str)
		return
	}
	fmt.Println(result)
}
