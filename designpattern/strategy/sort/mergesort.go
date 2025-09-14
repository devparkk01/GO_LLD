package main

import "fmt"

type MergeSort struct {
}

func (m * MergeSort) sort(arr []int) {
	fmt.Println("sorting using merge sort")
	m.mergeSortUtil(arr, 0 , len(arr) - 1)

}

func (m *MergeSort) mergeSortUtil(arr []int, left, right int ) {
	if left < right {
		mid := left + (right - left) / 2 
		m.mergeSortUtil(arr, left, mid )
		m.mergeSortUtil(arr, mid + 1,  right)
		m.merge(arr, left, right, mid)
	}
}


func (m *MergeSort) merge(arr []int, left, right, mid int ) {
	temp := make([]int , right - left + 1)

	i := left 
	j := mid + 1 
	k := 0 
	for(i <=mid && j <= right) {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			k++
			i++ 
		} else {
			temp[k] = arr[j]
			k++
			j++
		}
	}

	for(i <= mid) {
		temp[k] = arr[i] 
		k++ 
		i++ 
	}

	for( j <= right) {
		temp[k] = arr[j]
		j++ 
		k++ 
	}

	// temp array is sorted now 
	k = 0
	i = left 
	for(i <= right) {
		arr[i] = temp[k]
		i++ 
		k++ 
	}

}