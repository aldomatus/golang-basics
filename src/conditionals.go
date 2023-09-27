package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// conditionals
	var value1 bool = true
	var value2 bool = true

	if value1 == value2 {
		fmt.Printf("%t is equal to %t \n", value1, value2)
	}

	value1 = false

	if value1 != value2 {
		fmt.Printf("%t is different from %t \n", value1, value2)
	}

	numericValue1 := 1
	numericValue2 := 5

	if numericValue1 < numericValue2 {
		fmt.Printf("%d is greater than %d \n", numericValue1, numericValue2)
	}

	numericValue1 = 10
	numericValue2 = 5

	if numericValue1 > numericValue2 {
		fmt.Printf("%d is greater than %d \n", numericValue1, numericValue2)
	}

	numericValue1 = 5

	if numericValue1 >= numericValue2 {
		fmt.Printf("%d is greater than or equal to %d \n", numericValue1, numericValue2)
	}

	numericValue2 = 10

	if numericValue1 <= numericValue2 {
		fmt.Printf("%d is less than or equal to %d \n", numericValue1, numericValue2)
	}

	numericValue3 := 1
	numericValue4 := 2

	if numericValue1 < numericValue2 && numericValue3 < numericValue4 {
		fmt.Printf("%d is less than %d and %d is less than %d \n", numericValue1, numericValue2, numericValue3, numericValue4)
	}

	if numericValue1 > numericValue2 || numericValue3 < numericValue4 {
		fmt.Printf("%d is greater than %d or %d is less than %d \n", numericValue1, numericValue2, numericValue3, numericValue4)
	}

	if !value1 {
		fmt.Printf("!%t is true \n", value1)
	}

	// Using else if
	if value1 {
		fmt.Printf("!%t is true \n", value1)
	} else if value1 == false {
		fmt.Println("We are in else if")
	} else {
		fmt.Println("We are in else")
	}

	// Using else
	if value1 {
		fmt.Printf("!%t is true \n", value1)
	} else if value1 == true {
		fmt.Println("We are in else if")
	} else {
		fmt.Println("We are in else")
	}

	// convert string to number
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value:", value)
}
