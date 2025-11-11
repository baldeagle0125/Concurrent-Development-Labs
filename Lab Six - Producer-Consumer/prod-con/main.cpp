#include "SafeBuffer.h"
#include "Event.h"
#include "Semaphore.h"
#include <iostream>
#include <thread>
#include <vector>
#include <memory>
#include <mutex>


static const int num_threads = 100;
const int size = 20;
const int numLoops = 10;

// Mutex for console output to prevent garbled messages
std::mutex cout_mutex;


/*! \fn producer
    \brief Creates events and adds them to buffer
*/
void producer(std::shared_ptr<SafeBuffer<std::shared_ptr<Event>>> theBuffer, int numLoops, int id) {
  for(int i = 0; i < numLoops; ++i) {
    // Produce event and add to buffer
    std::shared_ptr<Event> e = std::make_shared<Event>(id * 1000 + i);
    theBuffer->put(e);
    {
      std::lock_guard<std::mutex> lock(cout_mutex);
      std::cout << "Producer " << id << " produced event " << e->getId() << std::endl;
    }
  }
}

/*! \fn consumer
    \brief Takes events from buffer and consumes them
*/
void consumer(std::shared_ptr<SafeBuffer<std::shared_ptr<Event>>> theBuffer, int numLoops, int id) {
  for(int i = 0; i < numLoops; ++i) {
    // Get event from buffer and consume it
    std::shared_ptr<Event> e = theBuffer->get();
    {
      std::lock_guard<std::mutex> lock(cout_mutex);
      std::cout << "Consumer " << id << " consuming event " << e->getId() << std::endl;
    }
  }
}

int main(void) {
  // Create shared buffer
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
  
  // Join all producer threads
  for(auto& t : producers) {
    t.join();
  }
  
  // Join all consumer threads
  for(auto& t : consumers) {
    t.join();
  }
  
  std::cout << "\nAll producers and consumers finished!" << std::endl;
  return 0;
}
