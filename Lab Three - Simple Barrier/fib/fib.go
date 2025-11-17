// Lab Three - Fibonacci Comparison (Sequential vs Parallel)
// Description: Compares sequential and parallel Fibonacci implementations
//              Demonstrates that fine-grained parallelism can be counterproductive

package main

import (
	"fmt"
	"sync"
)

// fib calculates Fibonacci number sequentially using recursion
// Parameters:
//   - N: Position in Fibonacci sequence
//
// Returns:
//   - Nth Fibonacci number
//
// Time Complexity: O(2^N) - exponential
func fib(N int) int {
	if N < 2 {
		return 1
	} else {
		return fib(N-1) + fib(N-2)
	}
}

// parFib calculates Fibonacci number in parallel using goroutines
// WARNING: This is extremely inefficient due to goroutine overhead
//
//	Creates exponential number of goroutines
//
// Parameters:
//   - N: Position in Fibonacci sequence
//
// Returns:
//   - Nth Fibonacci number
//
// Note: Demonstrates that not all recursive problems benefit from parallelization
func parFib(N int) int {
	var wg sync.WaitGroup
	var A, B int // Store results from recursive calls

	if N < 2 {
		return 1
	} else {
		wg.Add(2) // Wait for both recursive calls

		// Launch goroutine for fib(N-1)
		go func(N int, Ans *int) {
			defer wg.Done()
			*Ans = parFib(N - 1)
		}(N, &A)

		// Launch goroutine for fib(N-2)
		go func(N int, Ans *int) {
			defer wg.Done()
			*Ans = parFib(N - 2)
		}(N, &B)

		wg.Wait()    // Wait for both to complete
		return A + B // Return sum
	}
}

// main compares sequential vs parallel Fibonacci performance
// Results show both produce same values but parallel is MUCH slower
func main() {
	fmt.Println("Comparing Sequential vs Parallel Fibonacci")
	fmt.Println("Sequential --- Parallel")
	fmt.Println("==============================")

	// Calculate fib(0), fib(5), fib(10), ..., fib(45)
	for i := range 10 {
		Seq := fib(i * 5)    // Sequential version
		par := parFib(i * 5) // Parallel version (very slow!)
		fmt.Println(Seq, "---", par)
	}

	fmt.Println("\nNote: Parallel version is much slower due to goroutine overhead")
}
