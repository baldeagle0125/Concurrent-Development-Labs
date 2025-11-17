# Lab Six - Producer-Consumer

## Overview
Implementation of the classic producer-consumer problem using a thread-safe circular buffer with semaphore synchronization in C++.

## GitHub Repository
[https://github.com/baldeagle0125/Concurrent-Development-Labs](https://github.com/baldeagle0125/Concurrent-Development-Labs)

## Problem Statement
Multiple producer threads generate events and place them in a shared buffer, while multiple consumer threads remove and process events from the buffer. The solution must:
1. Prevent race conditions on the buffer
2. Block producers when buffer is full
3. Block consumers when buffer is empty
4. Ensure thread safety

## Implementation Details

### SafeBuffer Class (`SafeBuffer.h`)

A thread-safe circular buffer template class using semaphores:

**Key Components:**
- `mutex`: Semaphore (initialized to 1) protects buffer access
- `spaces`: Semaphore (initialized to buffer size) counts available spaces
- `items`: Semaphore (initialized to 0) counts items in buffer
- `first`: Index of first item (for removal)
- `last`: Index where next item will be added
- Circular array for storage

**Put Operation:**
```cpp
void put(T item) {
    spaces->Wait();   // Wait for available space
    mutex->Wait();    // Lock buffer
    buffer[last] = item;
    last = (last + 1) % size;
    mutex->Signal();  // Unlock buffer
    items->Signal();  // Signal that item is available
}
```

**Get Operation:**
```cpp
T get() {
    items->Wait();    // Wait for item to be available
    mutex->Wait();    // Lock buffer
    T item = buffer[first];
    first = (first + 1) % size;
    mutex->Signal();  // Unlock buffer
    spaces->Signal(); // Signal that space is available
    return item;
}
```

### Main Program (`main.cpp`)

**Configuration:**
- 100 total threads (50 producers + 50 consumers)
- Buffer size: 20 events
- Each producer creates 10 events
- Each consumer processes 10 events

**Event Class** (`Event.h`):
- Simple class with unique ID
- Tracks which producer created the event

**Console Output Protection:**
- Uses `std::mutex` to prevent garbled output messages
- Separate from buffer synchronization

## How to Run

```bash
cd "Lab Six - Producer-Consumer/prod-con"
make
./prodcon
```

## Expected Output

```
Producer 0 produced event 0
Producer 4 produced event 4000
Consumer 0 consuming event 0
Consumer 1 consuming event 2000
Producer 1 produced event 1001
...
All producers and consumers finished!
```

Events are produced and consumed concurrently without race conditions.

## Key Concepts

1. **Producer-Consumer Pattern**: Classic concurrent design pattern
2. **Bounded Buffer**: Fixed-size shared buffer
3. **Circular Buffer**: Efficient wraparound using modulo
4. **Counting Semaphores**: Track available spaces and items
5. **Binary Semaphore (Mutex)**: Protect critical sections
6. **Deadlock Prevention**: Careful semaphore ordering
7. **Template Classes**: Generic buffer implementation

## Semaphore Usage

| Semaphore | Initial Value | Purpose |
|-----------|--------------|---------|
| `mutex` | 1 | Mutual exclusion for buffer access |
| `spaces` | buffer_size | Count available spaces |
| `items` | 0 | Count items in buffer |

## Why This Works

1. **Producers blocked when full**: `spaces->Wait()` blocks when spaces == 0
2. **Consumers blocked when empty**: `items->Wait()` blocks when items == 0
3. **No race conditions**: `mutex` ensures only one thread modifies buffer
4. **No deadlock**: Consistent ordering (spaces/items before mutex)

## Files
- `main.cpp` - Main program with producer and consumer threads
- `SafeBuffer.h` - Thread-safe circular buffer template class
- `Event.h` - Simple event class
- `Semaphore.h` / `Semaphore.cpp` - Custom semaphore implementation
- `README` - Original C++ readme

## Alternative Implementations

Could also be implemented using:
- Condition variables (like Go's `sync.Cond`)
- Monitors
- Channels (in Go)
- Concurrent queues

## Learning Outcomes
- Understanding producer-consumer synchronization
- Implementing thread-safe data structures
- Working with counting and binary semaphores
- Circular buffer implementation
- Template class design in C++
- Preventing race conditions and deadlock
