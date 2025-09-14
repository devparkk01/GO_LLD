package main 

import (
	"fmt"
)

type BubbleSort struct{

}

func (s *BubbleSort) sort(arr []int) {
	fmt.Println("sorting using Bubble sort")
	n := len(arr)
	for i := 0 ; i < n; i++ {
		for j := 0 ; j < n - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
}