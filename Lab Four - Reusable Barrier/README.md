# Lab Four - Reusable Barrier

## Overview
Implementation of reusable barriers that can be used multiple times in the same program. Two different implementations are provided: one using atomic operations and channels, another using condition variables with phase tracking.

## GitHub Repository
[https://github.com/baldeagle0125/Concurrent-Development-Labs](https://github.com/baldeagle0125/Concurrent-Development-Labs)

## Problem Statement
Unlike simple barriers (Lab Three), reusable barriers must be able to synchronize multiple phases of execution without resetting manually. The barrier must work correctly even when goroutines reach it at different times in different phases.

## Implementation Details

### 1. Atomic Barrier (`atomic-barrier/atomic-barrier.go`)

Uses atomic operations and unbuffered channels:

**Key Components:**
- `arrived`: Atomic counter (`atomic.Int32`)
- `theChan`: Unbuffered channel for signaling
- Counter is incremented atomically
- Counter is decremented after passing through for reusability

**Algorithm:**
```
1. Atomically increment arrival counter
2. If last to arrive:
   - Send signal through channel
   - Receive signal back (turnstile)
   Else:
   - Receive signal (wait)
   - Send signal forward (pass it on)
3. Atomically decrement counter (prepare for next use)
```

**Demonstrates 2 barrier phases**: Parts A→B and C→D

### 2. Struct Barrier (`struct-barrier/struct-barrier.go`)

Uses a custom barrier struct with condition variables:

**Key Components:**
- `barrier` struct with mutex, condition variable, counter, and phase
- Phase tracking prevents race conditions
- Clean object-oriented design

**Barrier Struct:**
```go
type barrier struct {
    theLock sync.Mutex
    cond    *sync.Cond
    total   int
    count   int
    phase   int  // Tracks which phase we're in
}
```

**Algorithm:**
```
1. Lock mutex and save current phase
2. Increment counter
3. If last to arrive:
   - Reset counter to 0
   - Increment phase number
   - Broadcast to all waiting goroutines
   Else:
   - Wait while phase hasn't changed
4. Unlock mutex
```

**Demonstrates 2 barrier phases**: Parts A→B and C→D

## How to Run

### Atomic Barrier
```bash
cd "Lab Four - Reusable Barrier/atomic-barrier"
go run atomic-barrier.go
```

### Struct Barrier
```bash
cd "Lab Four - Reusable Barrier/struct-barrier"
go run struct-barrier.go
```

## Expected Output

Both implementations show proper synchronization across two phases:
```
Part A [random order]
...
PartB [random order]
Part C [random order]
...
PartD [random order]
```

All Part A completes before any Part B.
All Part C completes before any Part D.

## Key Concepts

1. **Reusable Barriers**: Barriers that can be used multiple times
2. **Phase Tracking**: Using phase numbers to prevent race conditions
3. **Atomic Operations**: Lock-free counter manipulation
4. **Unbuffered Channels**: Synchronous communication in Go
5. **Condition Variables**: Efficient waiting with `sync.Cond`
6. **Turnstile Pattern**: Sequential passage through barrier
7. **Object-Oriented Design**: Encapsulating barrier logic in a struct

## Comparison

| Feature | Atomic Barrier | Struct Barrier |
|---------|---------------|----------------|
| Synchronization | Atomic ops + Channel | Mutex + Cond |
| Readability | More complex | Cleaner API |
| Reusability | Manual counter reset | Phase tracking |
| Performance | Lock-free counter | Lock-based |

## Files
- `atomic-barrier/atomic-barrier.go` - Atomic implementation
- `struct-barrier/struct-barrier.go` - Struct-based implementation

## Learning Outcomes
- Understanding reusable vs simple barriers
- Working with atomic operations in Go
- Implementing phase tracking for reusability
- Using condition variables effectively
- Designing reusable synchronization primitives
- Channel-based synchronization patterns
