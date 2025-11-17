# Lab Two - Rendezvous

## Overview
Implementation of the rendezvous synchronization pattern where multiple goroutines must wait for each other at a common point before any can proceed.

## GitHub Repository
[https://github.com/baldeagle0125/Concurrent-Development-Labs](https://github.com/baldeagle0125/Concurrent-Development-Labs)

## Problem Statement
The rendezvous pattern ensures that no thread/goroutine can proceed past a certain point until all threads/goroutines have reached that point. This is also known as a barrier synchronization.

## Implementation Details

### Go Implementation (`rendezvous.go`)
- Uses `sync.Cond` (condition variable) for barrier synchronization
- Uses `sync.Mutex` to protect shared counter
- 5 goroutines synchronize at the barrier

**Key Components:**
- `grCount`: Counter for goroutines that have reached the barrier
- `threadCount`: Total number of goroutines
- `cond`: Condition variable for signaling
- `mutex`: Protects the shared counter

### C++ Implementation (`labTwo/rendezvous.cpp`)
- Uses custom `Semaphore` class
- Implements turnstile pattern
- Similar logic to Go version but using semaphores

## How to Run

### Go Version
```bash
cd "Lab Two - Rendezvous"
go run rendezvous.go
```

### C++ Version
```bash
cd "Lab Two - Rendezvous/labTwo"
make
./rendezvous
```

## Expected Output
```
Part A 3
Part A 1
Part A 2
Part A 0
Part A 4
PartB 4
PartB 3
PartB 0
PartB 1
PartB 2
```

Note: All "Part A" messages appear before any "PartB" messages, demonstrating proper barrier synchronization.

## Key Concepts

1. **Rendezvous Pattern**: All threads wait for each other at a synchronization point
2. **Condition Variables**: Efficient waiting mechanism using `sync.Cond`
3. **Broadcast vs Signal**: Last goroutine broadcasts to wake all waiting goroutines
4. **Critical Sections**: Protecting shared state with mutex

## Algorithm
```
1. Execute Part A
2. Acquire mutex
3. Increment counter
4. If last to arrive:
   - Broadcast signal to all waiting goroutines
   Else:
   - Wait for broadcast signal
5. Release mutex
6. Execute Part B
```

## Learning Outcomes
- Understanding barrier synchronization patterns
- Working with condition variables in Go
- Coordinating multiple concurrent threads
- Difference between broadcast and signal operations
