# Lab Three - Simple Barrier

## Overview
Implementation of a simple (non-reusable) barrier using mutex and semaphore to synchronize multiple goroutines.

## GitHub Repository
[https://github.com/baldeagle0125/Concurrent-Development-Labs](https://github.com/baldeagle0125/Concurrent-Development-Labs)

## Problem Statement
Create a barrier that ensures all goroutines complete Part A before any goroutine proceeds to Part B. This is a one-time barrier that doesn't need to be reused.

## Implementation Details

### Barrier Implementation (`barrier/barrier.go`)
Uses a combination of mutex and weighted semaphore from `golang.org/x/sync/semaphore`:

**Key Components:**
- `barrierMutex`: Protects the shared counter
- `count`: Tracks how many goroutines have arrived at the barrier
- `barrierSem`: Weighted semaphore initialized to 0 (blocked)
- `ctx`: Context for semaphore operations

**Algorithm:**
1. Each goroutine completes Part A
2. Locks mutex and increments counter
3. Last goroutine to arrive signals the semaphore
4. All goroutines wait on semaphore
5. Turnstile pattern allows all to pass through
6. Proceed to Part B

### Parallel Fibonacci (`fib/fib.go`)
Demonstrates parallel vs sequential Fibonacci calculation:
- Sequential version using simple recursion
- Parallel version spawning goroutines for each recursive call

**Warning**: Parallel version creates exponential goroutines and becomes very slow for larger numbers due to overhead.

## How to Run

### Barrier Example
```bash
cd "Lab Three - Simple Barrier"
go run barrier/barrier.go
```

### Fibonacci Example
```bash
cd "Lab Three - Simple Barrier"
go run fib/fib.go
```

## Expected Output (Barrier)
```
Part A 5
Part A 2
Part A 9
Part A 0
Part A 6
Part A 8
Part A 3
Part A 4
Part A 7
Part A 1
PartB 5
PartB 2
PartB 9
PartB 0
PartB 6
PartB 3
PartB 1
PartB 4
PartB 7
PartB 8
```

All "Part A" statements complete before any "PartB" statements execute.

## Key Concepts

1. **Barrier Synchronization**: Coordinating multiple threads at a rendezvous point
2. **Weighted Semaphores**: Using `golang.org/x/sync/semaphore` for blocking
3. **Turnstile Pattern**: Allowing all waiting goroutines to pass through sequentially
4. **Critical Sections**: Protecting shared state with mutex
5. **Fine-grained Parallelism**: Understanding when parallelism adds overhead (fibonacci example)

## Files
- `barrier/barrier.go` - Simple barrier implementation
- `fib/fib.go` - Parallel vs sequential Fibonacci comparison
- `go.mod` - Module dependencies

## Learning Outcomes
- Implementing barrier synchronization from scratch
- Understanding semaphore acquire/release patterns
- Working with weighted semaphores
- Recognizing when parallelism is counterproductive
