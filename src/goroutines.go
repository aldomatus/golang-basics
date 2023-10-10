package main

import (
	"fmt"
	"sync"
)

func say(wg *sync.WaitGroup, texts ...string) {
	// The defer in the Go code is used to schedule an action
	// that should be executed just before a function finishes its execution
	defer wg.Done()
	for _, text := range texts {
		fmt.Println(text)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(4)
	go say(&wg, "Hello", "World!")
	go say(&wg, "Hello 2", "World 2!")
	go say(&wg, "Hello 3", "World 3!")

	go func(wg *sync.WaitGroup, text string) {
		defer wg.Done()
		fmt.Println(text)
	}(&wg, "Good Bye!")

	// wg.Wait() in main blocks the execution until all the goroutines
	//registered in wg have called wg.Done(), which signifies they have finished.
	wg.Wait()
}
