package main

import "fmt"

func main() {
	// variables declaration
	helloMessage := "Hello"
	worldMessage := "world"

	// Println
	fmt.Println(helloMessage, worldMessage)

	// Printf
	name := "Aldo"
	age := 26

	fmt.Printf("%s is %d years old\n", name, age)
	fmt.Printf("%v is %v years old\n", name, age) // if you dont know the data type you can use %v

	// Sprintf
	message := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println(message)

	// Data type
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("age: %T\n", age)
}
