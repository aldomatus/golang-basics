package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Use the current time's nanoseconds to ensure a
	// different seed each run

	rand.Seed(time.Now().UnixNano())
	begin := time.Now()

	fmt.Println("====== Construction begins (with Goroutines) ======")

	wg := &sync.WaitGroup{}

	wg.Add(3)

	go Worker("Bob", "Installing a door", wg)
	go Worker("Alice", "mowing the lawn", wg)
	go Worker("John", "Painting the walls", wg)

	wg.Wait()
	fmt.Printf("====== Construction ended in %.2f seconds ======\n", time.Since(begin).Seconds())

}

// Worker initiates a job and takes a random number of
// seconds between 3 and 5 to finish it

func Worker(name, task string, wg *sync.WaitGroup) {
	fmt.Printf("Worker %s began to work on %s\n", name, task)

	for i := 1; i <= 10; i++ {
		fmt.Printf("Worker %s is working (%d/%d)...\n", name, i, 10)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Printf("Worker %s finished working on %s \n", name, task)
	wg.Done()
}
