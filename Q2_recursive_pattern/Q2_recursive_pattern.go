package main

import "fmt"

// Function to print a pattern recursively based on the value of n
func recursivePattern(n int) {
	// Base case: if n is less than or equal to 0, return and stop the recursion
	if n <= 0 {
		return
	}
	
	// If n is 9, print 9 and then recursively call the function with 4
	if n == 9 {
		fmt.Println(9)
		recursivePattern(4) // Call with 4
	} 
	// If n is 4, print 4 and then recursively call the function with 2
	else if n == 4 {
		fmt.Println(4)
		recursivePattern(2) // Call with 2
	} 
	// If n is 2, print 2 and the recursion stops here
	else if n == 2 {
		fmt.Println(2)
	}
}

func main() {
	// Start the recursive pattern by calling with 9
	recursivePattern(9)
}
