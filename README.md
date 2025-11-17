# Concurrent Development Labs

This repository contains all labs for the Concurrent Development module, demonstrating various concurrency patterns and synchronization techniques in Go and C++.

## GitHub Repository
[https://github.com/baldeagle0125/Concurrent-Development-Labs](https://github.com/baldeagle0125/Concurrent-Development-Labs)

## Labs Overview

### Set-up Lab
Introduction to Go basics including factorial, fibonacci, and functional programming concepts.

### Lab Two - Rendezvous
Implementation of the rendezvous pattern using `sync.Cond` for thread synchronization.

### Lab Three - Simple Barrier
Simple barrier implementation using mutex and semaphore to synchronize goroutines.

### Lab Four - Reusable Barrier
Two implementations of reusable barriers:
- Atomic barrier using atomic operations and channels
- Struct barrier using condition variables with phase tracking

### Lab Five - Dining Philosophers
Classic dining philosophers problem with deadlock prevention using resource hierarchy:
- Go implementation using buffered channels
- C++ implementation using semaphores

### Lab Six - Producer-Consumer
Thread-safe producer-consumer pattern using circular buffer with semaphore synchronization.

### Go Concurrency Essentials Lab
Collection of essential concurrency patterns:
- Atomic operations
- Mutex synchronization
- Semaphore patterns
- Channel signaling
- Worker pools

## Building and Running

### Go Programs
```bash
cd [lab-directory]
go run [filename].go
```

### C++ Programs
```bash
cd [lab-directory]
make
./[executable-name]
```

## License
GNU General Public License v3.0
