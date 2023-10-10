package main

import "fmt"

type Shape interface {
	area() float64
}

type Circle struct {
	radio float64
}

func (c Circle) area() float64 {
	return c.radio * c.radio * 3.1416
}

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func printArea(s Shape) {
	fmt.Println(s.area())
}

func main() {
	square := Square{side: 2}
	printArea(square)

	circle := Circle{radio: 2}
	printArea(circle)
}
