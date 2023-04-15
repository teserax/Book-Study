//Вспомним что такое интерфейс.
//Задача для этого: Напишите интерфейс
//для вычисления объема трехмерных фигур,
//таких как кубы и сферы. В чем польза
//интерфейса (попробуй дать ответ)? (1 день)

package main

import (
	"fmt"
	"math"
)

type Figures interface {
	volumeCalculations()
}

type Cube struct {
	height int
}
type Sphere struct {
	radius float64
}
type Triangle struct {
	heights int
	length  int
}

func (c Cube) volumeCalculations() {
	v := c.height * c.height * c.height
	fmt.Printf("The volume of the Cube is: %d\n", v)
}
func (s Sphere) volumeCalculations() {
	v := 4 / 3 * math.Pi * s.radius * 2
	fmt.Printf("The volume of the Sphere is: %.2f\n", v)
}
func (t Triangle) volumeCalculations() {
	v := t.length * t.heights / 2
	fmt.Printf("The volume of the Triangle is: %d\n", v)
}

func main() {
	var ball Figures = Cube{3}
	var box Figures = Sphere{4}
	var triangular Figures = Triangle{
		heights: 4,
		length:  4,
	}
	ball.volumeCalculations()
	box.volumeCalculations()
	triangular.volumeCalculations()
}

//интерфейс дает гибкость коду и переиспользование (и я думаю функциональность для других программистов ...
//зная что есть интерфейс который имплементирует данный метод или методы можно использовать без необходимости погруженния в суть исполнения
// что я имею веду можно добавить еще несколько  фигур(других обьектов)
//или заменить на другие но благодаря единому интерфейсу используя один методот
//можно взаимодействовать с несколькоми обьектами...
