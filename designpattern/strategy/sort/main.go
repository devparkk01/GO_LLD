package main

import "fmt"

func main() {
	arr := []int{1, 4, 200, 16, 2, 76, -78}
	mergeSort := &MergeSort{}
	sorter := NewSorter(mergeSort)
	sorter.Sort(arr)
	fmt.Println(arr)

	arr = []int{10, 4, 33, 2343, 8903, 2, 0, -20}
	bubbleSort := &BubbleSort{}
	sorter.SetSortStrategy(bubbleSort)
	sorter.Sort(arr)
	fmt.Println(arr)

}