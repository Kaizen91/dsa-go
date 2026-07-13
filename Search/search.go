package search

func LinearSearch(target int, array []int) int {
	for i, v := range array {
		if v == target {
			return i
		}
	}
	return -1
}
