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

	var hi string

	var results chan string

	// mutually exclusively
	var mutex sync.Mutex

	wg.Add(2) // starts two go routines

	go func() {
		mutex.Lock()           // the code between lock/unlock is the critical section. two go routines musnt be there at the same time.
		fmt.Println("hi" + hi) // io so the go routine is parked
		//mutex.Unlock()
		for i := 1; i <= 10; i += 1 {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%d goodbye\n", i)
		}

		wg.Done()
	}()
	go func() {
		// please may I modify the 'hi' variable?
		mutex.Lock() // if someone has the mutex, it should park.
		hi = "-hi"
		//mutex.Unlock()
		//fmt.Println("hi" + hi)
		for i := 1; i <= 10; i += 1 {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%d hello\n", i)
		}
		wg.Done()
	}()

	wg.Wait() // it waits until counter reaches 0. wg.Done() decrements it
}
