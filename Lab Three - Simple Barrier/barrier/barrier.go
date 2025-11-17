//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Ihor Melashchenko
// Issues: Fixed - barrier now properly implemented
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// ==================== GLOBAL VARIABLES ====================
// Barrier synchronization primitives (shared across goroutines)
var (
	barrierMutex sync.Mutex          // Protects shared counter
	count        int                 // Tracks goroutines at barrier
	barrierSem   *semaphore.Weighted // Blocks goroutines until all arrive
	ctx          context.Context     // Context for semaphore operations
)

// ==========================================================

// doStuff demonstrates barrier synchronization using mutex and semaphore
// All goroutines must complete Part A before any can proceed to Part B
// Parameters:
//   - goNum: Unique identifier for this goroutine
//   - wg: WaitGroup to signal completion
//   - totalRoutines: Total number of goroutines to synchronize
//
// Returns:
//   - bool: Always true (success indicator)
func doStuff(goNum int, wg *sync.WaitGroup, totalRoutines int) bool {
	// Simulate work before barrier
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	// ==================== BARRIER IMPLEMENTATION ====================
	// Step 1: Safely increment arrival counter
	barrierMutex.Lock()
	count++
	lastToArrive := (count == totalRoutines)
	barrierMutex.Unlock()

	// Step 2: Last goroutine signals the semaphore
	if lastToArrive {
		barrierSem.Release(1) // Open the turnstile
	}

	// Step 3: Wait at the turnstile (all goroutines pass through here)
	barrierSem.Acquire(ctx, 1) // Wait for/take the token
	barrierSem.Release(1)      // Pass token to next goroutine (turnstile pattern)
	// ================================================================

	// All goroutines have passed the barrier
	fmt.Println("PartB", goNum)
	wg.Done() // Signal completion
	return true
}

// main sets up and executes the barrier demonstration
func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	// ==================== BARRIER INITIALIZATION ====================
	ctx = context.TODO()                  // Create context for semaphore
	count = 0                             // Initialize arrival counter
	barrierSem = semaphore.NewWeighted(1) // Create semaphore with capacity 1
	barrierSem.Acquire(ctx, 1)            // Initialize to 0 (blocked/closed)
	// ================================================================

	// Launch all goroutines
	for i := range totalRoutines {
		go doStuff(i, &wg, totalRoutines)
	}

	// Wait for all goroutines to complete
	wg.Wait()
}
