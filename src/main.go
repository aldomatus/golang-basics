package main

import "fmt"

func main() {
	// constant declaration
	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// integer variables declaration
	base := 12
	var height int = 14
	var area int
	fmt.Println(base, height, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area of square
	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("Area of a square:", squareArea)

}
