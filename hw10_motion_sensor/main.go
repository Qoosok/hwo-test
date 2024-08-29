// fixed homework
package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"
)

func sensor(dataChan chan<- int) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var sensorValue int32
		err := binary.Read(rand.Reader, binary.LittleEndian, &sensorValue)
		if err != nil {
			continue
		}
		dataChan <- int(sensorValue % 100)
	}
}

func process(dataChan <-chan int, processedChan chan<- float64) {
	values := make([]int, 0, 10)

	for value := range dataChan {
		values = append(values, value)

		if len(values) == 10 {
			mean := calculateMean(values)
			processedChan <- mean
			values = values[:0] // Reset the slice
		}
	}
}

func calculateMean(values []int) float64 {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return float64(sum) / float64(len(values))
}

func main() {
	dataChan := make(chan int)
	processedChan := make(chan float64)

	go sensor(dataChan)
	go process(dataChan, processedChan)

	go func() {
		time.Sleep(time.Minute)
		close(dataChan)
	}()

	for mean := range processedChan {
		fmt.Printf("Mean of last 10 values: %.2f\n", mean)
	}
}
