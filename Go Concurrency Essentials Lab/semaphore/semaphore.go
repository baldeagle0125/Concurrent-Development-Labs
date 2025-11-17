// Go Concurrency Essentials - Semaphore Pattern Using Buffered Channel
// Description: Demonstrates using buffered channels as counting semaphores
//              to limit concurrent goroutine execution

package main

import (
	"fmt"
	"sync"
	"time"
)

// main demonstrates resource pool pattern using buffered channel as semaphore
func main() {
	maxGoroutines := 5 // Maximum concurrent goroutines allowed

	// Buffered channel acts as a counting semaphore
	// Capacity = max concurrent goroutines
	semaphore := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup

	// Launch 20 tasks, but only 5 can run concurrently
	for i := range 20 {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			// ==================== ACQUIRE SEMAPHORE ====================
			// Send to channel (blocks if channel is full)
			semaphore <- struct{}{}
			// ===========================================================

			// Ensure semaphore is released when goroutine exits
			defer func() {
				<-semaphore // Receive from channel (frees one slot)
			}()

			// ==================== CRITICAL WORK ====================
			// Simulate a task that takes 2 seconds
			fmt.Printf("Running task %d\n", i)
			time.Sleep(2 * time.Second)
			// =======================================================
		}(i)
	}

	// Wait for all tasks to complete
	wg.Wait()

	fmt.Println("\nAll tasks completed!")
}
