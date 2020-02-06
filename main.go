package main

import (
	"fmt"
	"runtime"
)

func main() {
	numThreads := runtime.NumCPU()
	fmt.Printf("I have : %d threads\n", numThreads)
	ch := make(chan int)
	for i := 0; i < numThreads; i++ {
		go sumOf3n5Multiples(1000, numThreads, 3+i, ch)
	}
	j := 0
	for i := 0; i < 8; i++ {
		j += <-ch

	}
	fmt.Printf("%d", j)
}

func sumOf3n5Multiples(maxNumber int, stepSize int, start int, ch chan int) {
	var sum = 0
	for i := start; i < maxNumber; i += stepSize {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	ch <- sum
}
