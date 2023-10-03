package main

import "fmt"

func main() {
	// Arrays
	var myArray [4]int
	fmt.Println(myArray)

	myArray[0] = 1
	myArray[1] = 2
	fmt.Println(myArray)

	fmt.Printf("len(myArray): %d\n", len(myArray))
	fmt.Printf("cap(myArray): %d\n", cap(myArray))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("slice: %v, len(slice): %d, cap(slice): %d\n", slice, len(slice), cap(slice))

	// Slice methods
	fmt.Println("slice[0]:", slice[0])
	fmt.Println("slice[:3]:", slice[:3])
	fmt.Println("slice[2:4]:", slice[2:4])
	fmt.Println("slice[4:]:", slice[4:])

	// Append
	slice = append(slice, 6)
	fmt.Println(slice)

	// Append new list
	newSlice := []int{7, 8, 9}
	slice = append(slice, newSlice...)
	fmt.Println("slice", slice)
}
