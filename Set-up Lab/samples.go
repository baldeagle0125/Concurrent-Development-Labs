// Set-up Lab - Introduction to Go Programming
// Description: Demonstrates basic Go concepts including functions, recursion,
//              higher-order functions, and lambda expressions

package main

import (
	"fmt"
)

// factorial calculates the factorial of a given number using iteration
// Parameters:
//   - N: The number to calculate factorial for
//
// Returns:
//   - The factorial of N (N!)
func factorial(N int64) int64 {
	var ans int64
	ans = 1

	// Multiply all numbers from 1 to N-1
	for i := int64(1); i < N; i++ {
		ans = ans * int64(i)
	}

	// Multiply by N to get final result
	return ans * N
}

// fib calculates the Nth Fibonacci number using recursion
// Parameters:
//   - N: The position in the Fibonacci sequence
//
// Returns:
//   - The Nth Fibonacci number
//
// Note: This is inefficient for large N due to exponential time complexity
func fib(N int) int {
	var ans int

	// Base cases: fib(1) = fib(2) = 1
	if N < 3 {
		ans = 1
	} else {
		// Recursive case: fib(n) = fib(n-1) + fib(n-2)
		ans = fib(N-1) + fib(N-2)
	}

	return ans
}

// applyMap applies a function to every element in a slice
// This demonstrates higher-order functions (functions as parameters)
// Parameters:
//   - theFun: The function to apply to each element
//   - theArray: The slice to transform (modified in-place)
func applyMap(theFun func(int) int, theArray []int) {
	for i := range len(theArray) {
		theArray[i] = theFun(theArray[i])
	}
}

// main demonstrates various Go programming concepts
func main() {
	var num int64
	num = 1

	// Interactive loop: calculate factorial and fibonacci for user input
	// Loop terminates when user enters 0
	for num != 0 {
		fmt.Print("Enter a number:")
		fmt.Scanln(&num)
		result := factorial(num)
		fmt.Println("the factorial is: ", result)
		fmt.Println("the fibonacci is:", fib(int(num)))
	}

	// Demonstrate lambda function (anonymous function)
	addOne := func(N int) int {
		return N + 1
	}

	// Use the lambda function directly
	X := addOne(9)
	fmt.Println(X)

	// Create and initialize a slice
	myArray := make([]int, 4)
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
	fmt.Println(myArray)

	// Apply the lambda function to all elements using higher-order function
	applyMap(addOne, myArray)
	fmt.Println(myArray)
}
