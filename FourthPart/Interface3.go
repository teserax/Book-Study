package main

import (
	"fmt"
)

type sum interface {
	Fold() string
	String() string
}

type myTemperatur struct {
	celsius int
}
type myString struct {
	name string
}

func (t myTemperatur) Fold() string {
	tem := t.celsius + t.celsius
	return fmt.Sprintf("sum of temperatures is:%v\n", tem)
}
func (t myTemperatur) String() string {
	return fmt.Sprintf("temperature is:%v", t.celsius)
}
func (s myString) Fold() string {
	return fmt.Sprintf("Hello %v\n", s.name)
}
func (s myString) String() string {
	return fmt.Sprintf("Name of user is:%v\n", s.name)
}
func test(s sum) func() string {

	t := s.Fold
	return t

}
func main() {
	var temp sum = myTemperatur{2}
	var greetings sum = myString{"Valera"}
	greetings.Fold()
	test(temp)
}
