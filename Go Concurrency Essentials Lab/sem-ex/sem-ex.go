// Go Concurrency Essentials - Weighted Semaphore Example
// Description: Demonstrates worker pool pattern using weighted semaphores
//              from golang.org/x/sync/semaphore package
//
// This example computes Collatz conjecture steps for numbers 1-64
// using a limited number of concurrent workers

package main

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"golang.org/x/sync/semaphore"
)

// main demonstrates worker pool using weighted semaphore
// Pattern: Limit concurrent goroutines to number of CPU cores
func main() {
	ctx := context.TODO()

	var (
		// Limit workers to number of available CPU cores
		maxWorkers = runtime.GOMAXPROCS(0)
		// Weighted semaphore with capacity = maxWorkers
		sem = semaphore.NewWeighted(int64(maxWorkers))
		// Output array for results
		out = make([]int, 64)
	)

	fmt.Printf("Computing with up to %d concurrent workers\n", maxWorkers)

	// Compute the output using up to maxWorkers goroutines at a time
	for i := range out {
		// ==================== ACQUIRE SEMAPHORE ====================
		// When maxWorkers goroutines are in flight, Acquire blocks
		// until one of the workers finishes
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}
		// ===========================================================

		// Launch worker goroutine
		go func(i int) {
			// Ensure semaphore is released when goroutine completes
			defer sem.Release(1)

			// Compute Collatz steps for this number
			out[i] = collatzSteps(i + 1)
		}(i)
	}

	// ==================== WAIT FOR COMPLETION ====================
	// Acquire all tokens to wait for any remaining workers to finish
	// Alternative: Use errgroup.Group for more sophisticated coordination
	if err := sem.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("Failed to acquire semaphore: %v", err)
	}
	// =============================================================

	fmt.Println("\nCollatz steps for numbers 1-64:")
	fmt.Println(out)
}

// collatzSteps computes the number of steps to reach 1 under the Collatz conjecture
// Collatz conjecture: For any positive integer n:
// - If n is even: n → n/2
// - If n is odd: n → 3n+1
// Repeat until reaching 1
//
// Parameters:
//   - n: Starting number (must be positive)
//
// Returns:
//   - Number of steps to reach 1
//
// Reference: https://en.wikipedia.org/wiki/Collatz_conjecture
func collatzSteps(n int) (steps int) {
	if n <= 0 {
		panic("nonpositive input")
	}

	for ; n > 1; steps++ {
		// Check for overflow (too many steps)
		if steps < 0 {
			panic("too many steps")
		}

		if n%2 == 0 {
			// Even: divide by 2
			n /= 2
			continue
		}

		// Odd: multiply by 3 and add 1
		// Check for integer overflow before computing
		const maxInt = int(^uint(0) >> 1)
		if n > (maxInt-1)/3 {
			panic("overflow")
		}
		n = 3*n + 1
	}

	return steps
}
