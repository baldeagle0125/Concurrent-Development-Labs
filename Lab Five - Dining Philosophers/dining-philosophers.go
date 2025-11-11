// Dining Philosophers Template Code
// Author: Joseph Kehoe
// Created: 21/10/24
//GPL Licence
// MISSING:
// 1. Readme
// 2. Full licence info.
// 3. Comments
// FIXED: Deadlock prevention using resource hierarchy solution

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was eating")
}

// Deadlock prevention: Last philosopher picks up forks in reverse order
func getForks(index int, forks map[int]chan bool, philCount int) {
	if index == philCount-1 {
		// Last philosopher picks up right fork first
		forks[(index+1)%philCount] <- true
		forks[index] <- true
	} else {
		// All other philosophers pick up left fork first
		forks[index] <- true
		forks[(index+1)%philCount] <- true
	}
}

func putForks(index int, forks map[int]chan bool, philCount int) {
	if index == philCount-1 {
		// Last philosopher releases in same order they acquired
		<-forks[(index+1)%philCount]
		<-forks[index]
	} else {
		<-forks[index]
		<-forks[(index+1)%philCount]
	}
}

func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool, philCount int, iterations int) {
	for range iterations {
		think(index)
		getForks(index, forks, philCount)
		eat(index)
		putForks(index, forks, philCount)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	philCount := 5
	iterations := 5 // Number of times each philosopher eats
	wg.Add(philCount)

	forks := make(map[int]chan bool)
	for k := range philCount {
		forks[k] = make(chan bool, 1)
	} //set up forks

	fmt.Println("Starting Dining Philosophers - Deadlock prevented using resource hierarchy")

	for N := range philCount {
		go doPhilStuff(N, &wg, forks, philCount, iterations)
	} //start philosophers

	wg.Wait() //wait here until everyone is done
	fmt.Println("All philosophers have finished dining!")
} //main
