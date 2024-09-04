package main

import (
	"sync"
	"testing"
)

func executeWorkers(increments []int) int {
	counter = 0
	counterMux = sync.Mutex{}
	wg = sync.WaitGroup{}

	numWorkers := len(increments)
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i, increments[i])
	}

	wg.Wait()
	return counter
}

func TestWorker(t *testing.T) {
	increment := 10

	increments := []int{increment}
	result := executeWorkers(increments)

	if result != increment {
		t.Errorf("Expected counter value %d, got %d", increment, result)
	}
}

func TestMainFunction(t *testing.T) {
	increments := []int{1, 2, 3, 4, 5}
	expected := 0
	for _, inc := range increments {
		expected += inc
	}

	result := executeWorkers(increments)

	if result != expected {
		t.Errorf("Expected counter value %d, got %d", expected, result)
	}
}

func BenchmarkWorker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		increments := []int{10}
		executeWorkers(increments)
	}
}
