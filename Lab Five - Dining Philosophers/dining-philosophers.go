// Lab Five - Dining Philosophers Problem
// Description: Classic dining philosophers problem with deadlock prevention
//              using resource hierarchy solution

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// think simulates a philosopher thinking for a random duration
// Parameters:
//   - index: Philosopher number (for output)
func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second)
	fmt.Println("Phil: ", index, "was thinking")
}

// eat simulates a philosopher eating for a random duration
// Parameters:
//   - index: Philosopher number (for output)
func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second)
	fmt.Println("Phil: ", index, "was eating")
}

// getForks acquires both forks for a philosopher
// Uses resource hierarchy to prevent deadlock:
// - Last philosopher picks up forks in reverse order
// Parameters:
//   - index: Philosopher number
//   - forks: Map of fork channels (buffered channels act as locks)
//   - philCount: Total number of philosophers
func getForks(index int, forks map[int]chan bool, philCount int) {
	if index == philCount-1 {
		// Last philosopher: RIGHT fork first (breaks circular wait)
		forks[(index+1)%philCount] <- true
		forks[index] <- true
	} else {
		// All other philosophers: LEFT fork first
		forks[index] <- true
		forks[(index+1)%philCount] <- true
	}
}

// putForks releases both forks after eating
// Must release in same order as acquisition
// Parameters:
//   - index: Philosopher number
//   - forks: Map of fork channels
//   - philCount: Total number of philosophers
func putForks(index int, forks map[int]chan bool, philCount int) {
	if index == philCount-1 {
		// Last philosopher: release in same order acquired
		<-forks[(index+1)%philCount]
		<-forks[index]
	} else {
		// Other philosophers: release in same order acquired
		<-forks[index]
		<-forks[(index+1)%philCount]
	}
}

// doPhilStuff simulates a philosopher's lifecycle
// Parameters:
//   - index: Philosopher number
//   - wg: WaitGroup to signal completion
//   - forks: Shared fork channels
//   - philCount: Total number of philosophers
//   - iterations: Number of eat-think cycles
func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool, philCount int, iterations int) {
	for range iterations {
		think(index)
		getForks(index, forks, philCount)
		eat(index)
		putForks(index, forks, philCount)
	}
	wg.Done() // Signal completion
}

// main sets up and runs the dining philosophers simulation
func main() {
	var wg sync.WaitGroup
	philCount := 5
	iterations := 5 // Number of times each philosopher eats
	wg.Add(philCount)

	// Initialize forks as buffered channels (capacity 1 = binary semaphore)
	forks := make(map[int]chan bool)
	for k := range philCount {
		forks[k] = make(chan bool, 1)
	}

	fmt.Println("Starting Dining Philosophers - Deadlock prevented using resource hierarchy")

	// Start all philosopher goroutines
	for N := range philCount {
		go doPhilStuff(N, &wg, forks, philCount, iterations)
	}

	wg.Wait() // Wait for all philosophers to finish
	fmt.Println("All philosophers have finished dining!")
}
