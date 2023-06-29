package main

import (
	"fmt"
	"time"
)

func hello() {
	for i := 0; i < 10; i += 1 {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d hello\n", i)
	}
}

func goodbye() {
	for i := 0; i < 10; i += 1 {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d goodbye\n", i)
	}
}

// (ready queue) main -> hello (for 1 sec) (print is io) -> main
func main() {
	go hello()
	go goodbye()
}
