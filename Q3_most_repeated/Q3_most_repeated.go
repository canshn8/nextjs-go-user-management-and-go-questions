package main

import "fmt"

// Function to find the most repeated string in the given array
func mostRepeated(arr []string) string {
	// Create a map to store the count of each string
	count := make(map[string]int)
	maxCount := 0    // Variable to track the maximum count of a string
	var result string // Variable to store the result (most repeated string)

	// Loop through each item in the array
	for _, item := range arr {
		// Increment the count of the current string in the map
		count[item]++
		
		// If the current string has the highest count, update the result and maxCount
		if count[item] > maxCount {
			maxCount = count[item] // Update maxCount with the new highest count
			result = item          // Update result with the current string
		}
	}

	return result
}

func main() {
	// Sample data array
	data := []string{"apple", "pie", "apple", "red", "red", "red"}

	// Print the most repeated string in the array
	fmt.Println(mostRepeated(data))
}
