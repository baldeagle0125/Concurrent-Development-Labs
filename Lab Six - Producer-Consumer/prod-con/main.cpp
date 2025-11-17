/**
 * Lab Six - Producer-Consumer Problem
 * Description: Demonstrates producer-consumer pattern using thread-safe
 *              circular buffer with semaphore synchronization
 * 
 * Configuration:
 * - 100 threads (50 producers + 50 consumers)
 * - Buffer capacity: 20 events
 * - Each producer creates 10 events
 * - Each consumer processes 10 events
 */

#include "SafeBuffer.h"
#include "Event.h"
#include "Semaphore.h"
#include <iostream>
#include <thread>
#include <vector>
#include <memory>
#include <mutex>

// ==================== CONFIGURATION ====================
static const int num_threads = 100;  // Total threads (producers + consumers)
const int size = 20;                 // Buffer capacity
const int numLoops = 10;             // Items per producer/consumer

// Mutex for console output to prevent interleaved messages
std::mutex cout_mutex;
// =======================================================

/**
 * producer - Creates events and adds them to the shared buffer
 * @param theBuffer: Shared thread-safe buffer
 * @param numLoops: Number of events to produce
 * @param id: Unique producer identifier
 */
void producer(std::shared_ptr<SafeBuffer<std::shared_ptr<Event>>> theBuffer, int numLoops, int id) {
  for(int i = 0; i < numLoops; ++i) {
    // Create new event with unique ID (id * 1000 + i)
    std::shared_ptr<Event> e = std::make_shared<Event>(id * 1000 + i);
    
    // Add to buffer (blocks if buffer is full)
    theBuffer->put(e);
    
    // Thread-safe console output
    {
      std::lock_guard<std::mutex> lock(cout_mutex);
      std::cout << "Producer " << id << " produced event " << e->getId() << std::endl;
    }
  }
}

/**
 * consumer - Takes events from buffer and processes them
 * @param theBuffer: Shared thread-safe buffer
 * @param numLoops: Number of events to consume
 * @param id: Unique consumer identifier
 */
void consumer(std::shared_ptr<SafeBuffer<std::shared_ptr<Event>>> theBuffer, int numLoops, int id) {
  for(int i = 0; i < numLoops; ++i) {
    // Get event from buffer (blocks if buffer is empty)
    std::shared_ptr<Event> e = theBuffer->get();
    
    // Thread-safe console output
    {
      std::lock_guard<std::mutex> lock(cout_mutex);
      std::cout << "Consumer " << id << " consuming event " << e->getId() << std::endl;
    }
  }
}

/**
 * main - Sets up and runs the producer-consumer simulation
 */
int main(void) {
  // Create shared buffer with capacity for 'size' events
  std::shared_ptr<SafeBuffer<std::shared_ptr<Event>>> aBuffer = 
    std::make_shared<SafeBuffer<std::shared_ptr<Event>>>(size);
  
  std::vector<std::thread> producers;
  std::vector<std::thread> consumers;
  
  // Create producer threads (half of total threads)
  for(int i = 0; i < num_threads/2; ++i) {
    producers.push_back(std::thread(producer, aBuffer, numLoops, i));
  }
  
  // Create consumer threads (half of total threads)
  for(int i = 0; i < num_threads/2; ++i) {
    consumers.push_back(std::thread(consumer, aBuffer, numLoops, i));
  }
  
  // Wait for all producer threads to complete
  for(auto& t : producers) {
    t.join();
  }
  
  // Wait for all consumer threads to complete
  for(auto& t : consumers) {
    t.join();
  }
  
  std::cout << "\nAll producers and consumers finished!" << std::endl;
  return 0;
}
