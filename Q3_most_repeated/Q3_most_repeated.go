package main

import "fmt"

func mostRepeated(arr []string) string {
	count := make(map[string]int)
	maxCount := 0
	var result string

	for _, item := range arr {
		count[item]++
		if count[item] > maxCount {
			maxCount = count[item]
			result = item
		}
	}

	return result
}

func main() {
	data := []string{"apple", "pie", "apple", "red", "red", "red"}
	fmt.Println(mostRepeated(data))
}