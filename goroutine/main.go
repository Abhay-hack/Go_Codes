package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine

func worker(id int) {
	fmt.Printf("Worker %d Starting\n", id)
	time.Sleep(time.Second) // sleeps for 1 swcond
	fmt.Printf("Worker %d Done\n", id)
}

func workeri(id int, wg *sync.WaitGroup) {
	defer wg.Done() // mark worker as done when it is finished
	fmt.Printf("Worker %d Starting\n", id)
	time.Sleep(time.Second) // sleeps for 1 swcond
	fmt.Printf("Worker %d Done\n", id)
}

func main() {
	// Concurrency with Routines and Channels

	// => GO routines are lightweight threads managed by go runtime, thousands of go routines can run concurrently

	for k := 1; k <= 3; k++ {
		go worker(k) // start workers concurrently
		// Each time it is creating new goroutine running worker(i). so it means 3 workers running at same time
	}
	time.Sleep(2 * time.Second) // wait for go routine to finish
	// => The main function itself is also a goroutine (the “main goroutine”). If main ends, the whole program ends—even if other goroutines are still running. So we force main to pause for 2 seconds to give workers enough time to finish (they each take 1s)

	//=> as time.sleep is a shortcut , the right way is to use sync.WaitGroup
	fmt.Println("Using sync.WaitGroup")
	var wg sync.WaitGroup
	for j := 1; j <= 3; j++ {
		wg.Add(1) //  expect one more worker
		go workeri(j, &wg)
	}
	wg.Wait() // wait for all workers to finish
	fmt.Println("All Workers Completed")

	// Channels
	// => safe way for goroutines to communicate. We use this syntax <- to send and recieve messages
	ch := make(chan string)

	// send msg from a goroutine
	go func() {
		ch <- "Message from goRoutine"
	}()
	// Recieve Message
	msg := <-ch
	fmt.Println(msg)

}
