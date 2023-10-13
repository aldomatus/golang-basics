package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Message1"
	c <- "Message2"

	fmt.Println(len(c), cap(c))

	// Range and close

	close(c) // close the channel
	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("Message1", email1)
	go message("Message2", email2)
	go message("Message3", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1: // waiting for a message through email1 channel
			fmt.Println("Email recived from email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recived from email2", m2)
		}
	}

}
