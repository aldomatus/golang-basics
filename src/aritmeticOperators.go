package main

import (
	"fmt"
	"math"
)

func main() {
	x := 10
	y := 50

	// Sum
	result := x + y
	fmt.Println("Sum:", result)

	// Subtract
	result = y - x
	fmt.Println("Subtract:", result)

	// Multiplication
	result = x * y
	fmt.Println("Multiplication:", result)

	// Division
	result = y / x
	fmt.Println("Division:", result)

	// Module
	result = y % x
	fmt.Println("Module:", result)

	// Incremental
	x++
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

	// Challenge: calculate the area of a
	// rectangle, trapeze and a circle
	fmt.Printf("Area of the rectangle: %d x %d = %d\n", x, y, x*y)

	height := 1
	fmt.Printf("Area of the trapeze: 0.5 *(%d + %d)*(%d) = %.2f\n", x, y, height, 0.5*float64(x+y)*float64(height))

	radio := 2
	fmt.Printf("Area of the circle: pi*(%d)^2 = %.2f\n", radio, math.Pi*float64(radio)*float64(radio))
}
