package main

import (
	"fmt"
	"sync"
	"time"
)

// data race - multiple concurrent at least 1 read and 1 write accesses to a variable. just read its fine.
// try -race with go tests too.

// (ready queue) main -> hello (for 1 sec) (print is io) -> main
func main() {
	var wg sync.WaitGroup

	// send to it or receive from it
	results := make(chan string)

	wg.Add(2) // starts two go routines

	go func() {
		for i := 1; i <= 10; i += 1 {
			time.Sleep(100 * time.Millisecond)
			// blocked on sending
			results <- fmt.Sprintf("%d goodbye", i)
		}

		wg.Done()
	}()
	go func() {
		for i := 1; i <= 10; i += 1 {
			time.Sleep(100 * time.Millisecond)
			// sends info / blocked on sending.
			results <- fmt.Sprintf("%d hello", i)
		}
		wg.Done()
	}()

	for result := range results {
		fmt.Println(result)
	}

	wg.Wait() // it waits until counter reaches 0. wg.Done() decrements it
}
