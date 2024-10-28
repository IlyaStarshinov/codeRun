package main

import (
	"fmt"
	"log"
)

func unicElem(arr []int, n int) int {
	count := 0
	for i := 0; i < n; {
		j := i + 1
		for j < n && arr[j] == arr[i] {
			j++
		}

		if j == i+1 {
			count++
		}

		i = j
	}
	return count
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	quickSort(arr, 0, n-1)
	fmt.Println(unicElem(arr, n))
}
