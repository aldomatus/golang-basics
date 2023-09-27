package main

import "fmt"

func printGreet(message string) {
	fmt.Println(message)
}

func tripleArguments(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	message := "Hello World"
	printGreet(message)

	tripleArguments(1, 2, "Hello")

	value := returnValue(2)
	fmt.Printf("Value: %d\n", value)

	value1, value2 := doubleReturn(2)
	fmt.Printf("%d and %d were returned\n", value1, value2)

	// If you wont use the second value you can replace it by "_"
	value_1, _ := doubleReturn(2)
	fmt.Printf("%d was returned\n", value_1)
}
