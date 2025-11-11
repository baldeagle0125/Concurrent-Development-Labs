#include "Semaphore.h"
#include <thread>
#include <vector>
#include <iostream>

/*! displays the first function in the barrier being executed */
void task(std::shared_ptr<Semaphore> mutexSem,std::shared_ptr<Semaphore> barrierSem, int threadCount){
  // Shared counter for threads that have reached the barrier
  static int count = 0; 
  
  std::cout << "first " << std::endl;
  
  //Barrier code - Rendezvous pattern

  // Lock the mutex to safely increment counter
  mutexSem->Wait();
  // Increment the count of threads that reached the barrier
  count++;
  // Unlock the mutex
  mutexSem->Signal(); 
  
  if(count == threadCount){
    // Last thread to arrive signals all waiting threads
    barrierSem->Signal();
  }
  
  // Wait for the signal from the last thread
  barrierSem->Wait(); 
  // Pass the signal to other waiting threads (turnstile pattern)
  barrierSem->Signal(); 
  
  std::cout << "second" << std::endl;
}

int main(void){
  std::shared_ptr<Semaphore> mutexSem;
  std::shared_ptr<Semaphore> barrierSem;
  int threadCount = 5;
  mutexSem=std::make_shared<Semaphore>(1);
  barrierSem=std::make_shared<Semaphore>(0);
  /*!< An array of threads*/
  std::vector<std::thread> threadArray(threadCount);
  /*!< Pointer to barrier object*/

  for(int i=0; i < threadArray.size(); i++){
    threadArray[i]=std::thread(task,mutexSem,barrierSem,threadCount);
  }

  for(int i = 0; i < threadArray.size(); i++){
    threadArray[i].join();
  }
  
  return 0;
}
