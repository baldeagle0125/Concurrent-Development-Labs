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

/*! \class SafeBuffer
    \brief A thread-safe circular buffer using semaphores
*/
template <typename T>
class SafeBuffer {
private:
    T* buffer;
    int size;
    int first;
    int last;
    std::shared_ptr<Semaphore> mutex;
    std::shared_ptr<Semaphore> spaces;  // Counts available spaces
    std::shared_ptr<Semaphore> items;   // Counts items in buffer
    
public:
    SafeBuffer(int s) : size(s), first(0), last(0) {
        buffer = new T[size];
        mutex = std::make_shared<Semaphore>(1);
        spaces = std::make_shared<Semaphore>(size);  // Initially all spaces available
        items = std::make_shared<Semaphore>(0);      // Initially no items
    }
    
    ~SafeBuffer() {
        delete[] buffer;
    }
    
    void put(T item) {
        spaces->Wait();  // Wait for available space
        mutex->Wait();   // Lock buffer
        
        buffer[last] = item;
        last = (last + 1) % size;
        
        mutex->Signal(); // Unlock buffer
        items->Signal(); // Signal that item is available
    }
    
    T get() {
        items->Wait();   // Wait for item to be available
        mutex->Wait();   // Lock buffer
        
        T item = buffer[first];
        first = (first + 1) % size;
        
        mutex->Signal(); // Unlock buffer
        spaces->Signal();// Signal that space is available
        
        return item;
    }
};

#endif

/* SafeBuffer.h ends here */
