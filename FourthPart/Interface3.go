package main

import "fmt"

type sum interface {
	Fold()
	String() string
}

type myTemperatur struct {
	celsius int
}
type myString struct {
	name string
}

func (t myTemperatur) Fold() {
	tem := t.celsius + t.celsius
	fmt.Printf("sum of temperatures is:%v\n", tem)
}
func (t myTemperatur) String() string {
	return fmt.Sprintf("temperature is:%v", t.celsius)
}
func (s myString) Fold() {
	fmt.Printf("Hello %v\n", s.name)
}
func (s myString) String() string {
	return fmt.Sprintf("Name of user is:%v\n", s.name)
}
func main() {
	var temp sum = myTemperatur{2}
	var greetings sum = myString{"Valera"}
	temp.Fold()
	fmt.Println(temp)
	fmt.Println("hello")
	greetings.Fold()
	fmt.Println(greetings)
}
