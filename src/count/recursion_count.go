package main

import "fmt"

func count(array []int) int {
	if len(array) == 0 {
		return 0
	}
	return 1 + count(array[1:])
}

func main() {
	mas := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, -6}
	fmt.Printf("%d", count(mas))
}
