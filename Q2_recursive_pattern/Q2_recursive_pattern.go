package main

import "fmt"

func recursivePattern(n int) {
	if n <= 0 {
		return
	}
	if n == 9 {
		fmt.Println(9)
		recursivePattern(4)
	} else if n == 4 {
		fmt.Println(4)
		recursivePattern(2)
	} else if n == 2 {
		fmt.Println(2)
	}
}

func main() {
	recursivePattern(9)
}