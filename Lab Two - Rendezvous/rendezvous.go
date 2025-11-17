// Lab Two - Rendezvous Pattern Implementation
// Description: Demonstrates the rendezvous synchronization pattern using sync.Cond
//              All goroutines must reach the barrier before any can proceed

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Global variables shared between goroutines
// Note: Using global variables is not considered an ideal practice
var (
	mutex       sync.Mutex // Protects access to grCount
	grCount     int        // Counts goroutines that have reached the barrier
	threadCount int        // Total number of goroutines
	cond        *sync.Cond // Condition variable for barrier synchronization
)

// WorkWithRendezvous demonstrates the rendezvous pattern
// Parameters:
//   - wg: WaitGroup to signal completion
//   - Num: Goroutine identifier for output
//
// Returns:
//   - bool: Always true (success indicator)
func WorkWithRendezvous(wg *sync.WaitGroup, Num int) bool {
	// Simulate random work duration (0-4 seconds)
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second)
	fmt.Println("Part A", Num)

	// ==================== RENDEZVOUS BARRIER ====================
	// This section ensures all goroutines wait for each other
	mutex.Lock()
	grCount++ // Increment count of arrived goroutines

	if grCount == threadCount {
		// Last goroutine to arrive: wake all waiting goroutines
		cond.Broadcast()
	} else {
		// Not all here yet: wait for signal from last goroutine
		cond.Wait() // Automatically releases and reacquires mutex
	}
	mutex.Unlock()
	// ============================================================

	// All goroutines have passed the barrier
	fmt.Println("PartB", Num)
	wg.Done() // Signal completion to WaitGroup
	return true
}

// main sets up and runs the rendezvous demonstration
func main() {
	var wg sync.WaitGroup
	threadCount = 5             // Number of goroutines to synchronize
	cond = sync.NewCond(&mutex) // Initialize condition variable with mutex

	wg.Add(threadCount)
	// Launch all goroutines
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N)
	}

	wg.Wait() // Wait for all goroutines to complete
}
