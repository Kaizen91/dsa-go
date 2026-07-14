package search

func LinearSearch(target int, array []int) int {
	for i, v := range array {
		if v == target {
			return i
		}
	}
	return -1
}

func BinarySearchRecursive(target int, array []int) int {
	if len(array) == 0 {
		return -1
	}
	mid := len(array) / 2
	if array[mid] == target {
		return mid
	} else if target < array[mid] {
		return BinarySearchRecursive(target, array[:mid])
	} else {
		subIndex := BinarySearchRecursive(target, array[mid+1:])
		// needed for search on the right side to not loose our place
		if subIndex != -1 {
			return mid + 1 + subIndex
		}
	}
	return -1
}

func BinarySearch(target int, array []int) int {
	hi := len(array)
	lo := 0
	for hi != lo {
		mid := (lo + hi) / 2
		if array[mid] == target {
			return mid
		} else if target > array[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return -1
}
