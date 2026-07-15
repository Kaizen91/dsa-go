package sort

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
