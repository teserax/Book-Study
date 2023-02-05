package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	total := 0
	if len(os.Args) == 1 {
		fmt.Println("please give me one argument!")
	} else {
		for i := 1; i < len(os.Args); i++ {
			number, err := strconv.Atoi(os.Args[i])
			if err != nil {
				fmt.Println(err)
			}
			total += number
		}
	}
	fmt.Println(total)
}
