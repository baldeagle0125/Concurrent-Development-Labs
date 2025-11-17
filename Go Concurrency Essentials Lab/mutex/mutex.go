// Go Concurrency Essentials - Mutex Example
// Description: Demonstrates traditional mutex-based synchronization
//              for protecting shared state

package main

import (
	"fmt"
	"sync"
)

// Global variables (passed by reference in practice)
var wg sync.WaitGroup
var total int64 // Shared counter (not thread-safe without mutex)

// adds increments a counter in a loop with mutex protection
// Parameters:
//   - n: Number of times to increment
//   - theLock: Mutex to protect shared counter
//
// Returns:
//   - bool: Always true (success indicator)
func adds(n int, theLock *sync.Mutex) bool {
	for range n {
		// ==================== CRITICAL SECTION ====================
		theLock.Lock()   // Acquire exclusive access
		total++          // Increment shared variable
		theLock.Unlock() // Release lock
		// ==========================================================
	}
	wg.Done() // Signal completion to WaitGroup
	return true
}

// main demonstrates mutex synchronization with multiple goroutines
// Result: 10 goroutines × 1000 increments each = 10,000 (always correct)
func main() {
	// Mutex passed by reference (better than global, though still used here for demo)
	var theLock sync.Mutex

	total = 0
	wg.Add(10) // Initialize WaitGroup for 10 goroutines

	// Launch 10 goroutines, each incrementing 1000 times
	for i := range 10 {
		fmt.Println(i)
		go adds(1000, &theLock)
	}

	wg.Wait() // Wait for all goroutines to complete

	// Print final value (should always be 10,000)
	fmt.Println(total)
}
