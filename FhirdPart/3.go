package main

import (
	"fmt"
	"strings"
	"time"
)

func timeParce(str string) {

	strTime := strings.Split(strings.ReplaceAll(str, "Published", ""), ",")
	for _, val := range strTime {
		tt := strings.TrimSpace(val)
		t, err := time.Parse("2 January 2006", tt)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(t.Format("2006-01-02"))
	}

}
func main() {
	str := "Published 25 March 2022,Published 8 March 2023"
	timeParce(str)
}
