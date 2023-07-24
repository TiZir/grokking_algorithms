package main

import "fmt"

func max(array []int) int {
	if len(array) == 2 {
		if array[0] > array[1] {
			return array[0]
		} else {
			return array[1]
		}
	}
	maxElem := max(array[1:])
	if array[0] > maxElem {
		return array[0]
	} else {
		return maxElem
	}
}

func main() {
	mas := []int{-7, -2, -3, -45, -5, -6, -1, 9, -9, -10, -11, 8}
	fmt.Printf("%d", max(mas))
}
