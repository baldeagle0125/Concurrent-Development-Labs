// Go Concurrency Essentials - Atomic Operations Example
// Description: Demonstrates lock-free synchronization using atomic operations
//              for simple counter increments

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Global WaitGroup for barrier synchronization
var wg sync.WaitGroup

// addsAtomic increments an atomic counter in a loop
// Uses atomic.Add for thread-safe increments without locks
// Parameters:
//   - n: Number of times to increment
//   - total: Atomic counter (shared between goroutines)
//
// Returns:
//   - bool: Always true (success indicator)
func addsAtomic(n int, total *atomic.Int64) bool {
	for range n {
		total.Add(1) // Atomically increment by 1
	}
	wg.Done() // Signal completion to WaitGroup
	return true
}

// main demonstrates atomic operations with multiple goroutines
// Result: 10 goroutines × 1000 increments each = 10,000 (always correct)
func main() {
	var total atomic.Int64 // Atomic 64-bit integer

	// Launch 10 goroutines, each incrementing 1000 times
	for i := range 10 {
		wg.Add(1) // Add to WaitGroup before starting goroutine
		fmt.Println("go Routine ", i)
		go addsAtomic(1000, &total)
	}

	wg.Wait() // Wait for all goroutines to complete

	// Load and print final value (should always be 10,000)
	fmt.Println(total.Load())
}
