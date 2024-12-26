package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortWordsByA(words []string) []string {
	sort.SliceStable(words, func(i, j int) bool {
		countA_i := strings.Count(words[i], "a")
		countA_j := strings.Count(words[j], "a")

		if countA_i == countA_j {
			return len(words[i]) > len(words[j])
		}
		return countA_i > countA_j
	})

	return words
}

func main() {
	words := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	fmt.Println(sortWordsByA(words))
}