package main

import (
	"fmt"
	"log"
)

func bubbleSort(arr []int) int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr[1]
}

func main() {
	var arr = make([]int, 3)
	for i := 0; i < 3; i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Print(bubbleSort(arr))
}
