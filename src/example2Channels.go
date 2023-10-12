// Pool of workers

/*
This Go program demonstrates the use of channels and goroutines to create a worker pool.

1. Start of the Program: Two channels named 'jobs' and 'results' and a sync.WaitGroup named 'wg' are created.

2. Launching Workers: Three anonymous goroutines are launched, each executing the function 'worker'. The WaitGroup is notified of a new "task" using 'wg.Add(1)'.

3. Sending Jobs: Jobs are sent to the 'jobs' channel from 1 to 9 using a for loop.

4. Close Jobs: The 'jobs' channel is closed after all the jobs have been sent.

5. Wait and Close 'results': Another goroutine waits for all the workers to finish (using 'wg.Wait()') before closing the 'results' channel.

6. Collect Results: Finally, the results from the 'results' channel are collected and printed.

Output Observation:
The reason "worker 1" performs the majority of the jobs might be because its goroutine was the fastest to pick up jobs from the 'jobs' channel.
This is influenced by the Go runtime and the OS, which determine how tasks are allocated to CPUs. Moreover, the channel access is not predefined
in a specific order for the goroutines; they pick values from the channel as soon as they can.

*/

package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "processing job:", job)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results)
		}(w)
	}

	// Send jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("result:", result)
	}
}
