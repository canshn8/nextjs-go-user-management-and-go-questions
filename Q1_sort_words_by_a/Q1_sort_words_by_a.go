package main

import (
	"fmt"
	"sort"
	"strings"
)

// Function to sort words by the number of 'a' characters and then by length
func sortWordsByA(words []string) []string {
	// Sorting the slice of words using custom criteria
	sort.SliceStable(words, func(i, j int) bool {
		// Count 'a' characters in the ith and jth words
		countA_i := strings.Count(words[i], "a")
		countA_j := strings.Count(words[j], "a")

		// If the counts are equal, compare the lengths of the words
		if countA_i == countA_j {
			return len(words[i]) > len(words[j]) // Longer words come first
		}
		// Otherwise, sort by the number of 'a' characters in descending order
		return countA_i > countA_j
	})

	return words
}

func main() {
	// Sample list of words
	words := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}

	// Print the sorted list based on the number of 'a' characters and word length
	fmt.Println(sortWordsByA(words))
}
