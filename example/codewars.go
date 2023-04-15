package main

import (
	"fmt"
)

func FakeBin(x string) string {
	// your code here
	str := ""
	for _, val := range x {
		s := string(val)
		if s >= "5" {
			str += "1"
		} else {
			str += "0"
		}
	}
	return str
}

func main() {
	fmt.Println(FakeBin("254"))

}
