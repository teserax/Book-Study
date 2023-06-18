package main

import "fmt"

func main() {
	romNum := map[string]int{}
	romNum["I"] = 1
	romNum["V"] = 5
	romNum["X"] = 10
	romNum["L"] = 50
	romNum["C"] = 100
	romNum["D"] = 500
	romNum["M"] = 1000
	s := "MMMCMXCIX"

	result := 0
	test := 0
	for i := len(s) - 1; i >= 0; i-- {
		if _, ok := romNum[string(s[i])]; ok {
			if test > romNum[string(s[i])] {
				result -= romNum[string(s[i])]
			} else {
				result += romNum[string(s[i])]
				test = romNum[string(s[i])]
			}

		}
	}
	fmt.Println(result)
}
