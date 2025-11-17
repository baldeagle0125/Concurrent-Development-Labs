//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Ihor Melashchenko
// Description:
// A reusable barrier implemented using atomic variable and unbuffered channel
// Issues: None
//1. Change mutex to atomic variable - DONE
//2. Make it a reusable barrier - DONE
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// doStuff demonstrates reusable barrier using atomic operations and channels
// Can be used multiple times in sequence (demonstrates reusability)
// Parameters:
//   - goNum: Unique identifier for this goroutine
//   - arrived: Atomic counter tracking goroutines at barrier
//   - max: Total number of goroutines to synchronize
//   - wg: WaitGroup to signal completion
//   - theChan: Unbuffered channel for turnstile pattern
//
// Returns:
//   - bool: Always true (success indicator)
func doStuff(goNum int, arrived *atomic.Int32, max int, wg *sync.WaitGroup, theChan chan bool) bool {
	// ==================== FIRST PHASE ====================
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	// Barrier 1: Using atomic operations and channel turnstile
	currentCount := arrived.Add(1) // Atomically increment counter

	if currentCount == int32(max) { // Last to arrive - signal others to go
		theChan <- true // Open turnstile
		<-theChan       // Close after going through
	} else { // Not all here yet - wait until signal
		<-theChan       // Wait at turnstile
		theChan <- true // Pass signal to next goroutine
	}

	// Reset for reusability - decrement counter atomically
	arrived.Add(-1)

	fmt.Println("PartB", goNum)

	// ==================== SECOND PHASE (Demonstrates Reusability) ====================
	time.Sleep(time.Second)
	fmt.Println("Part C", goNum)

	// Barrier 2: Same barrier, reused for second synchronization point
	currentCount = arrived.Add(1)
	if currentCount == int32(max) {
		theChan <- true
		<-theChan
	} else {
		<-theChan
		theChan <- true
	}
	arrived.Add(-1)

	fmt.Println("PartD", goNum)
	wg.Done() // Signal completion
	return true
}

// main sets up and executes the reusable barrier demonstration
func main() {
	totalRoutines := 10
	var arrived atomic.Int32 // Atomic counter (lock-free)
	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	// Unbuffered channel acts as turnstile for barrier
	theChan := make(chan bool)

	// Launch all goroutines
	for i := range totalRoutines {
		go doStuff(i, &arrived, totalRoutines, &wg, theChan)
	}

	// Wait for all goroutines to complete both phases
	wg.Wait()
}
