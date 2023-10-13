package main

import (
	"fmt"
	"time"
)

func urgentTask(c chan string) {
	time.Sleep(4 * time.Second)
	c <- "Urgent Task"
}

func importantTask(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "Important Task"
}

func normalTask(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "Normal Task"
}

func main() {
	urgentChan := make(chan string)
	importantChan := make(chan string)
	normalChan := make(chan string)

	go urgentTask(urgentChan)
	go importantTask(importantChan)
	go normalTask(normalChan)

	for i := 0; i < 3; i++ {
		select {
		case task := <-urgentChan:
			fmt.Println("Executing:", task)
		case task := <-importantChan:
			fmt.Println("Executing:", task)
		case task := <-normalChan:
			fmt.Println("Executing:", task)
		}
	}
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func urgentTask(c chan string) {
// 	time.Sleep(4 * time.Second)
// 	c <- "Urgent Task"
// }

// func importantTask(c chan string) {
// 	time.Sleep(2 * time.Second)
// 	c <- "Important Task"
// }

// func normalTask(c chan string) {
// 	time.Sleep(1 * time.Second)
// 	c <- "Normal Task"
// }

// func main() {
// 	urgentChan := make(chan string)
// 	importantChan := make(chan string)
// 	normalChan := make(chan string)

// 	go urgentTask(urgentChan)
// 	go importantTask(importantChan)
// 	go normalTask(normalChan)

// 	// This will block the main goroutine until the urgent task is done
// 	task := <-urgentChan
// 	fmt.Println("Executing:", task)

// 	// Now that we're unblocked, we can execute the other tasks
// 	task = <-importantChan
// 	fmt.Println("Executing:", task)

// 	task = <-normalChan
// 	fmt.Println("Executing:", task)
// }
