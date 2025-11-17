# Go Concurrency Essentials Lab

## Overview
Collection of essential Go concurrency patterns and synchronization techniques, demonstrating various approaches to concurrent programming in Go.

## GitHub Repository
[https://github.com/baldeagle0125/Concurrent-Development-Labs](https://github.com/baldeagle0125/Concurrent-Development-Labs)

## Examples Included

### 1. Atomic Operations (`atomic/atomic.go`)

Demonstrates atomic operations for lock-free synchronization:
- Uses `sync/atomic.Int64` for thread-safe counter
- 10 goroutines each increment counter 1000 times
- Final result: 10,000 (correct)
- No mutex needed for simple counters

**Key Concept**: Atomic operations are faster than mutex for simple operations.

### 2. Mutex Synchronization (`mutex/mutex.go`)

Traditional mutex-based synchronization:
- Uses `sync.Mutex` to protect shared counter
- 10 goroutines each increment counter 1000 times
- Demonstrates critical sections
- Final result: 10,000 (correct)

**Key Concept**: Mutexes provide mutual exclusion for complex critical sections.

### 3. Semaphore Pattern (`semaphore/semaphore.go`)

Buffered channel as counting semaphore:
- Limits concurrent goroutines to 5
- 20 tasks compete for 5 slots
- Demonstrates resource pool pattern
- Tasks wait when all 5 slots occupied

**Key Concept**: Buffered channels can implement semaphores elegantly.

### 4. Signalling Pattern (`signalling/signalling.go`)

Channel-based signalling for rendezvous:
- Two goroutines coordinate execution
- One goroutine waits for signal from other
- Demonstrates unbuffered channel communication
- Simple 2-thread rendezvous

**Key Concept**: Unbuffered channels provide synchronous communication.

### 5. Weighted Semaphore (`sem-ex/sem-ex.go`)

Worker pool using `golang.org/x/sync/semaphore`:
- Limits concurrent workers to `GOMAXPROCS`
- 64 tasks processed by limited worker pool
- Demonstrates weighted semaphore API
- Computes Collatz conjecture steps

**Key Concept**: Weighted semaphores allow acquiring multiple tokens.

## How to Run

```bash
# Atomic operations
cd "Go Concurrency Essentials Lab/atomic"
go run atomic.go

# Mutex synchronization
cd "Go Concurrency Essentials Lab/mutex"
go run mutex.go

# Semaphore pattern
cd "Go Concurrency Essentials Lab/semaphore"
go run semaphore.go

# Signalling pattern
cd "Go Concurrency Essentials Lab/signalling"
go run signalling.go

# Weighted semaphore
cd "Go Concurrency Essentials Lab/sem-ex"
go run sem-ex.go
```

## Expected Outputs

### Atomic
```
go Routine 0
...
go Routine 9
10000
```

### Mutex
```
0
1
...
9
10000
```

### Semaphore
```
Running task 4
Running task 9
...
(only 5 running at once)
```

### Signalling
```
StuffOne - Part A
StuffTwo - Part A
StuffTwo - PartB
StuffOne - PartB
```

### Sem-ex
```
[0 1 7 2 5 8 16 3 19 6 14 9 9 17 17 4 ...]
```

## Comparison of Techniques

| Technique | Use Case | Pros | Cons |
|-----------|----------|------|------|
| Atomic | Simple counters | Fast, lock-free | Limited operations |
| Mutex | Complex critical sections | Flexible | Slower, can deadlock |
| Buffered Channel | Resource pools | Idiomatic Go | Memory overhead |
| Unbuffered Channel | Signalling | Synchronous | Blocking |
| Weighted Semaphore | Worker pools | Fine control | External package |

## Key Concepts

1. **Atomic Operations**: Lock-free synchronization for simple types
2. **Mutual Exclusion**: Traditional locking with mutex
3. **Channels as Semaphores**: Idiomatic Go concurrency
4. **Unbuffered Channels**: Synchronous communication
5. **Weighted Semaphores**: Advanced resource management
6. **WaitGroups**: Barrier synchronization
7. **Worker Pools**: Limiting concurrency

## Best Practices Demonstrated

1. Use atomic operations for simple counters
2. Use mutexes for complex critical sections
3. Prefer channels for communication
4. Use buffered channels for resource pools
5. Always use `defer` with `WaitGroup.Done()`
6. Pass synchronization primitives by reference

## Learning Outcomes
- Understanding different synchronization primitives in Go
- Choosing appropriate technique for each scenario
- Working with channels, mutexes, and atomic operations
- Implementing common concurrency patterns
- Resource pool and worker pool patterns

## Files
- `atomic/atomic.go` - Atomic operations example
- `mutex/mutex.go` - Mutex synchronization example
- `semaphore/semaphore.go` - Buffered channel semaphore
- `signalling/signalling.go` - Channel signalling
- `sem-ex/sem-ex.go` - Weighted semaphore worker pool
- `sem-ex/go.mod` - Module dependencies
