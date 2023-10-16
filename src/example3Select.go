// If the Goroutine takes too long (more than 1 second, in this case) to send the message into the channel,
// the  will fall into the timeout case, and the program will print "timeout".

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
