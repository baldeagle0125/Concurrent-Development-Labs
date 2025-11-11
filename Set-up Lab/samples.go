package main

import (
	"fmt"
)

func factorial(N int64) int64 {
	var ans int64
	ans = 1

	for i := int64(1); i < N; i++ {
		ans = ans * int64(i)
	}

	return ans * N
}

func fib(N int) int {
	var ans int

	if N < 3 {
		ans = 1
	} else {
		ans = fib(N-1) + fib(N-2)
	}

	return ans
}

func applyMap(theFun func(int) int, theArray []int) {
	for i := range len(theArray) {
		theArray[i] = theFun(theArray[i])
	}
}

func main() {
	var num int64
	num = 1

	for num != 0 {
		fmt.Print("Enter a number:")
		fmt.Scanln(&num)
		result := factorial(num)
		fmt.Println("the factorial is: ", result)
		fmt.Println("the fibonacci is:", fib(int(num)))
	}

	addOne := func(N int) int {
		return N + 1
	}

	X := addOne(9)
	fmt.Println(X)
	myArray := make([]int, 4)
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
	fmt.Println(myArray)
	applyMap(addOne, myArray)
	fmt.Println(myArray)
}
