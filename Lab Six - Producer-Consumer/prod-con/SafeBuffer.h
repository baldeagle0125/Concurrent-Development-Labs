/* SafeBuffer.h --- 
 * 
 * Filename: SafeBuffer.h
 * Description: 
 * Author: Joseph
 * Maintainer: 
 * Created: Tue Jan  8 12:30:23 2019 (+0000)
 * Version: 
 * Package-Requires: ()
 * Last-Updated: Tue Jan  8 12:30:25 2019 (+0000)
 *           By: Joseph
 *     Update #: 1
 * URL: 
 * Doc URL: 
 * Keywords: 
 * Compatibility: 
 * 
 */

/* Commentary: 
 * 
 * 
 * 
 */

/* Change Log:
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

#ifndef SAFEBUFFER_H
#define SAFEBUFFER_H

#include "Semaphore.h"
#include <memory>
#include <mutex>

/**
 * SafeBuffer - A thread-safe circular buffer using semaphore synchronization
 * Template class that works with any data type
 * 
 * Thread Safety Strategy:
 * - mutex: Binary semaphore for mutual exclusion
 * - spaces: Counting semaphore tracking available buffer spaces
 * - items: Counting semaphore tracking items in buffer
 * 
 * Prevents:
 * - Race conditions (via mutex)
 * - Buffer overflow (via spaces semaphore)
 * - Buffer underflow (via items semaphore)
 */
template <typename T>
class SafeBuffer {
private:
    T* buffer;                              // Circular array storage
    int size;                               // Buffer capacity
    int first;                              // Index of first item (for get)
    int last;                               // Index where next item goes (for put)
    std::shared_ptr<Semaphore> mutex;       // Protects buffer operations
    std::shared_ptr<Semaphore> spaces;      // Counts available spaces
    std::shared_ptr<Semaphore> items;       // Counts items in buffer
    
public:
    /**
     * Constructor - Initializes buffer and semaphores
     * @param s: Buffer capacity (size)
     */
    SafeBuffer(int s) : size(s), first(0), last(0) {
        buffer = new T[size];
        mutex = std::make_shared<Semaphore>(1);     // Binary semaphore
        spaces = std::make_shared<Semaphore>(size); // All spaces initially free
        items = std::make_shared<Semaphore>(0);     // No items initially
    }
    
    /**
     * Destructor - Frees allocated buffer memory
     */
    ~SafeBuffer() {
        delete[] buffer;
    }
    
    /**
     * put - Adds an item to the buffer
     * Blocks if buffer is full
     * @param item: Item to add to buffer
     * 
     * Synchronization:
     * 1. Wait for available space (blocks if full)
     * 2. Lock buffer for exclusive access
     * 3. Add item and update index
     * 4. Unlock buffer
     * 5. Signal that new item is available
     */
    void put(T item) {
        spaces->Wait();  // Wait for available space (blocks if full)
        mutex->Wait();   // Lock buffer for exclusive access
        
        // Add item to circular buffer
        buffer[last] = item;
        last = (last + 1) % size;  // Wrap around using modulo
        
        mutex->Signal(); // Unlock buffer
        items->Signal(); // Signal that item is available for consumers
    }
    
    /**
     * get - Removes and returns an item from the buffer
     * Blocks if buffer is empty
     * @return: Item removed from buffer
     * 
     * Synchronization:
     * 1. Wait for item to be available (blocks if empty)
     * 2. Lock buffer for exclusive access
     * 3. Remove item and update index
     * 4. Unlock buffer
     * 5. Signal that space is now available
     */
    T get() {
        items->Wait();   // Wait for item to be available (blocks if empty)
        mutex->Wait();   // Lock buffer for exclusive access
        
        // Remove item from circular buffer
        T item = buffer[first];
        first = (first + 1) % size;  // Wrap around using modulo
        
        mutex->Signal(); // Unlock buffer
        spaces->Signal();// Signal that space is available for producers
        
        return item;
    }
};

#endif

/* SafeBuffer.h ends here */
