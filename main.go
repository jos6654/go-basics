package main

import (
	"fmt"
)

// Time and space complexity O(n)
func fib(num int) int {
	fibSlice := []int{1, 1}
	for i := 2; i <= num; i++ {
		fibSlice = append(fibSlice, fibSlice[i-1]+fibSlice[i-2])
	}
	return fibSlice[num-1]
}

func main() {
	var num int
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&num)
	fmt.Printf("The element at position %d of the Fibonnaci sequence is %d", num, fib(num))
}
