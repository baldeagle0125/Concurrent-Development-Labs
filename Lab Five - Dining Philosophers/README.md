# Lab Five - Dining Philosophers

## Overview
Implementation of the classic Dining Philosophers problem with deadlock prevention using the resource hierarchy solution. Both Go and C++ implementations are provided.

## GitHub Repository
[https://github.com/baldeagle0125/Concurrent-Development-Labs](https://github.com/baldeagle0125/Concurrent-Development-Labs)

## Problem Statement
Five philosophers sit at a round table with five forks between them. Each philosopher alternates between thinking and eating. To eat, a philosopher needs both adjacent forks. The challenge is to design a solution that:
1. Avoids deadlock
2. Avoids starvation
3. Allows maximum concurrency

## Deadlock Prevention Strategy

### Resource Hierarchy Solution
- Philosophers 0-3: Pick up **left fork first**, then right fork
- Philosopher 4 (last): Pick up **right fork first**, then left fork

This breaks the circular wait condition, preventing deadlock.

## Implementation Details

### Go Implementation (`dining-philosophers.go`)

**Key Components:**
- Buffered channels (`chan bool`) represent forks
- Each fork channel has capacity 1
- 5 philosophers, each eating 5 times
- Random sleep times for thinking and eating

**Fork Acquisition:**
```go
func getForks(index int, forks map[int]chan bool, philCount int) {
    if index == philCount-1 {
        // Last philosopher: right fork first
        forks[(index+1)%philCount] <- true
        forks[index] <- true
    } else {
        // Others: left fork first
        forks[index] <- true
        forks[(index+1)%philCount] <- true
    }
}
```

### C++ Implementation (`philosophers/main.cpp`)

**Key Components:**
- Vector of semaphores (`std::shared_ptr<Semaphore>`) represent forks
- Each semaphore initialized to 1 (available)
- Same resource hierarchy strategy
- Random sleep times using `sleep()`

**Files:**
- `main.cpp` - Main program logic
- `Semaphore.h` / `Semaphore.cpp` - Custom semaphore implementation

## How to Run

### Go Version
```bash
cd "Lab Five - Dining Philosophers"
go run dining-philosophers.go
```

### C++ Version
```bash
cd "Lab Five - Dining Philosophers/philosophers"
make
./philosophers
```

## Expected Output

Both versions show philosophers alternating between thinking and eating:
```
Starting Dining Philosophers - Deadlock prevented using resource hierarchy
Phil: 0 was thinking
Phil: 1 was thinking
...
Phil: 3 was eating
Phil: 2 was eating
...
All philosophers have finished dining!
```

No deadlock occurs, and all philosophers complete their meals.

## Key Concepts

1. **Dining Philosophers Problem**: Classic concurrency problem
2. **Deadlock Prevention**: Breaking circular wait condition
3. **Resource Hierarchy**: Ordering resources to prevent deadlock
4. **Starvation Avoidance**: Ensuring all philosophers get to eat
5. **Buffered Channels (Go)**: Using channels as locks
6. **Semaphores (C++)**: Binary semaphores as mutexes

## Deadlock Conditions (and how we prevent them)

| Condition | Prevention Strategy |
|-----------|---------------------|
| Mutual Exclusion | Cannot prevent (forks are exclusive) |
| Hold and Wait | Cannot easily prevent |
| No Preemption | Cannot prevent (can't steal forks) |
| **Circular Wait** | **Prevented by resource hierarchy** ✓ |

## Alternative Solutions (Not Implemented)

1. **Waiter Solution**: Central coordinator grants permission
2. **Chandy/Misra Solution**: Token-based, allows more concurrency
3. **Limiting Diners**: Allow only N-1 philosophers to eat simultaneously

## Files
- `dining-philosophers.go` - Go implementation
- `philosophers/main.cpp` - C++ main program
- `philosophers/Semaphore.h` - Semaphore header
- `philosophers/Semaphore.cpp` - Semaphore implementation
- `philosophers/README` - Original C++ readme

## Learning Outcomes
- Understanding classic concurrency problems
- Implementing deadlock prevention strategies
- Resource hierarchy ordering technique
- Working with semaphores and channels as synchronization primitives
- Comparing Go and C++ concurrency approaches
