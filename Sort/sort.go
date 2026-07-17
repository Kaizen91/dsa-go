package sort

import "fmt"

func BubbleSort(arr []int) {
	// compare two points if the first is bigger swap the positions
	// continue down the length of the array
	// then continue to l - i where l is length and i is the iteration
	l := len(arr)
	for i := 0; i < l ; i++{
		for j := 0; j < l - 1 - i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func partition(arr []int, hi, low int) int {
	// returns a partition  index 
	pivot := arr[hi]
	idx := low -1
	for i:=low; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}
	idx++
	arr[hi], arr[idx] = arr[idx], arr[hi]
	return idx
}

func qs(arr []int, hi, low int) {
	if low >= hi {
		return
	}
	//recurses
	pivot_idx := partition(arr, hi, low)
	qs(arr, pivot_idx - 1, low)
	qs(arr, hi, pivot_idx + 1)
}

func QuickSort(arr []int) {
	// entry point
	qs(arr, len(arr) - 1, 0)
}

func merge(arr []int, hi, mid, lo int) {
	L := make([]int, mid-lo)
	R := make([]int, hi-mid)
	copy(L, arr[lo:mid])
	copy(R, arr[mid:hi])
	l_idx, r_idx := 0, 0
	fmt.Printf("arr=%v, L=%v, R=%v\n", arr, L, R)
	for i:=lo; i < hi; i++ {
		fmt.Printf(
			"lo=%v, mid=%v, hi=%v, i=%d, l_idx=%d, r_idx=%d\n",
			lo, mid, hi, i, l_idx, r_idx,
		)
		if l_idx >= len(L) {
			arr[i] = R[r_idx]
			r_idx++
		} else if r_idx >= len(R) {
			arr[i] = L[l_idx]
			l_idx++
		} else if L[l_idx] <= R[r_idx] {
			arr[i] = L[l_idx]
			l_idx++
		} else {
			arr[i] = R[r_idx]
			r_idx++
		}
	}
}

func merge_sort_recursive(arr []int, hi, lo int) {
	if hi - lo > 1 {
		mid := (hi + lo) / 2
		merge_sort_recursive(arr, mid, lo)
		merge_sort_recursive(arr, hi, mid)
		merge(arr, hi, mid, lo)
	}
}

func MergeSort(arr []int) {
	hi, lo := len(arr), 0
	merge_sort_recursive(arr, hi, lo)
}
