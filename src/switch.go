package main

import "fmt"

func printDayOfWeek(number int) {
	switch number {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid number")
	}
}

func main() {

	switch module := 4 % 2; module {
	case 0:
		fmt.Println("It's even")
	default:
		fmt.Println("It's odd")
	}

	// Switch without condition
	value := 200

	switch {
	case value > 100:
		fmt.Printf("%d is greater than 100 \n", value)
	case value < 0:
		fmt.Printf("%d is less than 0 \n", value)
	default:
		fmt.Println("No condition")
	}

	printDayOfWeek(3) // Output will be "Tuesday"
	printDayOfWeek(8) // Output will be "Invalid number"
}
