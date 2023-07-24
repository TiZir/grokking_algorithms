package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
)

func binary_search(list [128]string, item string) (int, error) {
	left := 0
	right := 128
	for left <= right {
		mid := (left + right) / 2
		guess := list[mid]
		if guess == item {
			return mid, nil
		}
		if guess > item {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return 0, errors.New("can't found")
}

func main() {
	f, err := os.Open("names.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	os.Stdin = f
	var listNames [128]string
	for i := 0; i < 128; i++ {
		if _, err := fmt.Scanf("%s", &listNames[i]); err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
	}
	sort.Strings(listNames[:])
	result, err := binary_search(listNames, "Mia")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("%d", result)
}
