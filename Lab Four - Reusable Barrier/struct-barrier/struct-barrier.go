package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Create a reusable barrier data type
type barrier struct {
	theLock sync.Mutex
	cond    *sync.Cond
	total   int
	count   int
	phase   int // Track which phase we're in
}

// creates a properly initialised barrier
// N== number of threads (go Routines)
func createBarrier(N int) *barrier {
	theBarrier := &barrier{
		total: N,
		count: 0,
		phase: 0,
	}
	theBarrier.cond = sync.NewCond(&theBarrier.theLock)
	return theBarrier
}

// Method belonging to barrier data type
// Blocks until everyone reaches this point then lets everyone continue
// Reusable barrier implementation using condition variable
func (b *barrier) wait() {
	b.theLock.Lock()
	currentPhase := b.phase
	b.count++

	if b.count == b.total {
		// Last goroutine to arrive - wake everyone and prepare for next cycle
		b.count = 0
		b.phase++ // Move to next phase
		b.cond.Broadcast()
	} else {
		// Wait until all goroutines arrive (phase changes)
		for currentPhase == b.phase {
			b.cond.Wait()
		}
	}
	b.theLock.Unlock()
} //wait

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, theBarrier *barrier) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	//First Rendezvous here
	theBarrier.wait()

	fmt.Println("PartB", Num)

	// Demonstrate reusability - second barrier
	time.Sleep(time.Duration(rand.IntN(3)) * time.Second)
	fmt.Println("Part C", Num)
	theBarrier.wait()

	fmt.Println("PartD", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	barrier := createBarrier(5)
	threadCount := 5

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, barrier)
	}
	wg.Wait() //wait here until everyone (5 go routines) is done

}
