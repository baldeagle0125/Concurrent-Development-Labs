// Go Concurrency Essentials - Channel Signalling Example
// Description: Demonstrates unbuffered channel for goroutine synchronization
//              Simple 2-goroutine rendezvous pattern

package main

import (
	"fmt"
	"sync"
	"time"
)

// main demonstrates channel-based signalling between two goroutines
func main() {
	var wg sync.WaitGroup

	// Unbuffered channel for synchronization
	// Send blocks until receive happens (synchronous communication)
	barrier := make(chan bool)

	// doStuffOne executes first part, signals, then continues
	doStuffOne := func() bool {
		fmt.Println("StuffOne - Part A")

		// ==================== SEND SIGNAL ====================
		// Send signal (blocks until received by other goroutine)
		barrier <- true
		// =====================================================

		fmt.Println("StuffOne - PartB")
		wg.Done()
		return true
	}

	// doStuffTwo waits for signal before continuing
	doStuffTwo := func() bool {
		// Simulate some work (5 seconds)
		time.Sleep(time.Second * 5)
		fmt.Println("StuffTwo - Part A")

		// ==================== WAIT FOR SIGNAL ====================
		// Receive signal (blocks until sent by other goroutine)
		<-barrier
		// =========================================================

		fmt.Println("StuffTwo - PartB")
		wg.Done()
		return true
	}

	wg.Add(2)
	go doStuffOne()
	go doStuffTwo()

	// Wait for both goroutines to complete
	wg.Wait()

	fmt.Println("\nBoth goroutines synchronized successfully!")
}
