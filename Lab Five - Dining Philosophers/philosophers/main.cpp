/* main.c --- 
 * 
 * Filename: main.c
 * Description: 
 * Author: Joseph
 * Maintainer: 
 * Created: Wed Oct 11 09:28:12 2023 (+0100)
 * Last-Updated: Wed Oct 11 10:01:39 2023 (+0100)
 *           By: Joseph
 *     Update #: 13
 * 
 */

/* Commentary: 
 * 
 * 
 * 
 */

/* This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or (at
 * your option) any later version.
 * 
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * General Public License for more details.
 * 
 * You should have received a copy of the GNU General Public License
 * along with GNU Emacs.  If not, see <http://www.gnu.org/licenses/>.
 */

/* Code: */

#include "Semaphore.h"
#include <iostream>
#include <thread>
#include <vector>
#include <memory>
#include <stdlib.h>     /* srand, rand */
#include <time.h>       /* time */
#include<unistd.h>

const int COUNT = 5;
const int THINKTIME=3;
const int EATTIME=5;
const int ITERATIONS=5; // Number of times each philosopher eats
std::vector<std::shared_ptr<Semaphore>> forks(COUNT);


void think(int myID){
  int seconds=rand() % THINKTIME + 1;
  std::cout << myID << " is thinking! "<<std::endl;
  sleep(seconds);
}

// Deadlock prevention: Last philosopher picks up forks in reverse order
void get_forks(int philID){
  if (philID == COUNT - 1) {
    // Last philosopher picks up right fork first (reverse order)
    forks[(philID+1)%COUNT]->Wait();
    forks[philID]->Wait();
  } else {
    // All other philosophers pick up left fork first
    forks[philID]->Wait();
    forks[(philID+1)%COUNT]->Wait();
  }
}

void put_forks(int philID){
  if (philID == COUNT - 1) {
    // Last philosopher releases in reverse order
    forks[(philID+1)%COUNT]->Signal();
    forks[philID]->Signal();
  } else {
    forks[philID]->Signal();
    forks[(philID+1)%COUNT]->Signal();
  }
}

void eat(int myID){
  int seconds=rand() % EATTIME + 1;
    std::cout << myID << " is chomping! "<<std::endl;
  sleep(seconds);  
}

void philosopher(int id){
  for(int i = 0; i < ITERATIONS; i++){
    think(id);
    get_forks(id);
    eat(id);
    put_forks(id);
  }//for
  std::cout << "Philosopher " << id << " has finished dining!" << std::endl;
}//philosopher



int main(void){
  srand (time(NULL)); // initialize random seed:
  
  // Initialize all forks to 1 (available) using shared_ptr
  for(int i = 0; i < COUNT; i++){
    forks[i] = std::make_shared<Semaphore>(1);
  }
  
  std::cout << "Starting Dining Philosophers - Deadlock prevented using resource hierarchy" << std::endl;
  
  std::vector<std::thread> vt(COUNT);
  int id=0;
  for(std::thread& t: vt){
    t=std::thread(philosopher,id++);
  }
  /**< Join the philosopher threads with the main thread */
  for (auto& v :vt){
      v.join();
  }
  
  std::cout << "All philosophers have finished dining!" << std::endl;
  return 0;
}

/* main.c ends here */
