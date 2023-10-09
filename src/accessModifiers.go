package main

import (
	"fmt"
	pk "golang_course/src/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 1990
	fmt.Println(myCar)

	pk.PrintMessage("Public")
}
