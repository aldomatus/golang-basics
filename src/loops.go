package main

import "fmt"

func main() {
	// Conditional For
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	// For While
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever
	for {
		fmt.Println("loop")
		break
	}

	fmt.Printf("\n")

	// For inverse
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	// For While inverse
	counter = 10
	for counter > 0 {
		fmt.Println(counter)
		counter--
	}
}
