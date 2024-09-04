package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter    int
	counterMux sync.Mutex
	wg         sync.WaitGroup
)

func worker(id int, increment int) {
	defer wg.Done()

	time.Sleep(time.Millisecond * time.Duration(100+id*10))

	counterMux.Lock()
	counter += increment
	counterMux.Unlock()

	fmt.Printf("Worker %d completed incrementing by %d\n", id, increment)
}

func main() {
	numWorkers := 5
	increments := []int{1, 2, 3, 4, 5}

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, increments[i])
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}
