// Data Pipeline
package main

import (
	"fmt"
	"time"
)

// Simulate a long process by sleeping for 1 second
func longProcess(n int) int {
	time.Sleep(time.Second)
	return n
}

// First step: Multiply by 2
func multiplyByTwo(input <-chan int, stage1 chan<- int) {
	fmt.Println("multiplyByTwo")
	for number := range input {
		fmt.Printf("multiplying %d by 2\n", number)
		stage1 <- longProcess(number * 2)
	}
	close(stage1)
}

// Second step: add 3
func AddThree(input <-chan int, stage2 chan<- int) {
	fmt.Println("AddThree")
	for number := range input {
		fmt.Printf("Adding 3 to %d\n", number)
		stage2 <- longProcess(number + 3)
	}
	close(stage2)
}

// Third step: multiply by 4
func multiplyByFour(input <-chan int, output chan<- int) {
	fmt.Println("multiplyByFour")
	for number := range input {
		fmt.Printf("multiplying %d by 4\n", number)
		output <- longProcess(number) * 4
	}
	close(output)
}

func main() {
	input := make(chan int, 4)
	stage1 := make(chan int, 4)
	stage2 := make(chan int, 4)
	output := make(chan int, 4)

	for i := 1; i <= 4; i++ {
		input <- i
	}

	close(input)

	go multiplyByTwo(input, stage1)
	go AddThree(stage1, stage2)
	go multiplyByFour(stage2, output)

	for number := range output {
		fmt.Println(number)
	}

}
