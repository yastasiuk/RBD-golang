package main

import "fmt"

func main() {
	fmt.Println("FibonacciIterative: ", FibonacciIterative(5))
	fmt.Println("FibonacciRecursive: ", FibonacciRecursive(5))
	fmt.Println("IsPrime: ", IsPrime(101))
	fmt.Println("IsBinaryPalindrome: ", IsBinaryPalindrome(5))
	fmt.Println("ValidParentheses: ", ValidParentheses("[]"))
	fmt.Println("Increment: ", Increment("10"))
}
