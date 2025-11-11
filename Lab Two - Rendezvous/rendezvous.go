package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Global variables shared between functions --A BAD IDEA
var (
	mutex       sync.Mutex
	grCount     int
	threadCount int
	cond        *sync.Cond
)

func WorkWithRendezvous(wg *sync.WaitGroup, Num int) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)

	//Rendezvous here - using sync.Cond for barrier
	mutex.Lock()
	grCount++
	if grCount == threadCount {
		cond.Broadcast() // Last goroutine wakes all waiting goroutines
	} else {
		cond.Wait() // Wait until all goroutines reach this point
	}
	mutex.Unlock()

	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	threadCount = 5
	cond = sync.NewCond(&mutex)

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done

}
