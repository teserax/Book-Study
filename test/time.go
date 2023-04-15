package main

import (
	"fmt"
	"time"
)

func main() {
	d := time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC)
	d1 := time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(2000, 1, 3, 0, 0, 0, 0, time.UTC)
	// ... operation that takes 20 milliseconds ...
	elapsed := d.Sub(d1)
	fmt.Println(elapsed)
	elapsed = d.Sub(d2)
	fmt.Println(elapsed)

}
