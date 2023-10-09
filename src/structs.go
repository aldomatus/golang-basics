package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{
		brand: "Ford",
		year:  2020}
	fmt.Println(myCar)

	// another way to do it
	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.year = 2023

	fmt.Println(otherCar)
}
