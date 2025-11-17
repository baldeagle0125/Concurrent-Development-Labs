// Lab Four - Reusable Barrier (Struct-Based Implementation)
// Description: Implements a reusable barrier using a struct with phase tracking
//              Demonstrates object-oriented approach to synchronization primitives

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// ==================== BARRIER DATA TYPE ====================
// barrier is a reusable synchronization primitive that blocks goroutines
// until all have reached the barrier point
type barrier struct {
	theLock sync.Mutex // Protects shared state
	cond    *sync.Cond // Condition variable for signaling
	total   int        // Total number of goroutines to synchronize
	count   int        // Current number of arrived goroutines
	phase   int        // Current phase number (for reusability)
}

// ===========================================================

// createBarrier constructs and initializes a new barrier
// Parameters:
//   - N: Number of goroutines that must reach barrier before release
//
// Returns:
//   - Pointer to initialized barrier
func createBarrier(N int) *barrier {
	theBarrier := &barrier{
		total: N,
		count: 0,
		phase: 0,
	}
	// Bind condition variable to the barrier's mutex
	theBarrier.cond = sync.NewCond(&theBarrier.theLock)
	return theBarrier
}

// wait blocks until all goroutines reach the barrier
// Uses phase tracking to enable reusability
// This is a method on the barrier type (receiver: b *barrier)
func (b *barrier) wait() {
	b.theLock.Lock()
	currentPhase := b.phase // Remember which phase we entered in
	b.count++

	if b.count == b.total {
		// Last goroutine to arrive - wake everyone and prepare for next cycle
		b.count = 0        // Reset counter for next use
		b.phase++          // Move to next phase
		b.cond.Broadcast() // Wake all waiting goroutines
	} else {
		// Wait until all goroutines arrive (phase changes)
		// Keep waiting while still in same phase
		for currentPhase == b.phase {
			b.cond.Wait() // Releases lock while waiting, reacquires when signaled
		}
	}
	b.theLock.Unlock()
}

// WorkWithRendezvous demonstrates using the reusable barrier
// Parameters:
//   - wg: WaitGroup to signal completion
//   - Num: Goroutine identifier
//   - theBarrier: Shared barrier object
//
// Returns:
//   - bool: Always true (success indicator)
func WorkWithRendezvous(wg *sync.WaitGroup, Num int, theBarrier *barrier) bool {
	// ==================== FIRST PHASE ====================
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) // Random work duration
	fmt.Println("Part A", Num)

	// First Rendezvous: all goroutines wait here
	theBarrier.wait()

	fmt.Println("PartB", Num)

	// ==================== SECOND PHASE (Demonstrates Reusability) ====================
	time.Sleep(time.Duration(rand.IntN(3)) * time.Second)
	fmt.Println("Part C", Num)

	// Second Rendezvous: barrier reused for second synchronization point
	theBarrier.wait()

	fmt.Println("PartD", Num)
	wg.Done() // Signal completion
	return true
}

// main sets up and executes the reusable barrier demonstration
func main() {
	var wg sync.WaitGroup
	threadCount := 5

	// Create barrier for 5 goroutines
	barrier := createBarrier(threadCount)

	wg.Add(threadCount)
	// Launch all goroutines
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, barrier)
	}

	// Wait for all goroutines to complete both phases
	wg.Wait()
}
