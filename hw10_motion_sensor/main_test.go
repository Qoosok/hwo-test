package main

import (
	"testing"
)

func TestCalculateMean(t *testing.T) {
	values := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	expectedMean := 55.0

	mean := calculateMean(values)
	if mean != expectedMean {
		t.Errorf("Expected mean %.2f but got %.2f", expectedMean, mean)
	}
}

func TestProcess(t *testing.T) {
	dataChan := make(chan int)
	processedChan := make(chan float64)
	go process(dataChan, processedChan)

	go func() {
		values := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
		for _, v := range values {
			dataChan <- v
		}
		close(dataChan)
	}()

	expectedMean := 55.0
	mean := <-processedChan
	if mean != expectedMean {
		t.Errorf("Expected mean %.2f but got %.2f", expectedMean, mean)
	}
}
