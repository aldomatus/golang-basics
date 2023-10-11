package main

import "fmt"

func main() {
	conveyorBelt := make(chan string) // capacity 0

	go func() {
		fmt.Println("Sending piece 1 to conveyor belt")
		conveyorBelt <- "Piece 1"
		fmt.Println("Piece 1 sent")
	}()

	fmt.Println("Waiting piece 1 in conveyor belt...")
	piece1 := <-conveyorBelt
	fmt.Println("Piece 1 received:", piece1)

}
