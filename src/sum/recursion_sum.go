package main

import "fmt"

// func sum(array []int, i int) int {
// 	if i == len(array) {
// 		return 0
// 	}
// 	return array[i] + sum(array, i+1)
// }

//	func main() {
//		mas := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, -6}
//		fmt.Printf("%d", sum(mas, 0))
//	}
func sum(array []int) int {
	if len(array) == 0 {
		return 0
	}
	return array[0] + sum(array[1:])
}

func main() {
	mas := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, -6}
	fmt.Printf("%d", sum(mas))
}
