package main

import "fmt"

func main() {
	// Defer
	defer fmt.Println("hello")
	fmt.Println("world")

	for i := 0; i < 10; i++ {
		if i == 2 {
			fmt.Println("It's 2")
			continue
		}
		if i == 8 {
			fmt.Println("It's 8")
			break
		}
		fmt.Println(i)
	}

}
