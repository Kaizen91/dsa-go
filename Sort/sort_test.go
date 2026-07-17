package sort

import (
	"testing"
	"slices"
)

type SharedTestCase struct {
	name string
	arr []int
	expected []int
}

func GetSharedTestCases() []SharedTestCase {
	return []SharedTestCase{
		{
			name: "unsorted array",
			arr: []int{5,4,7,1,3,6,9,8,2},
			expected: []int{1,2,3,4,5,6,7,8,9},
		},
	}
}

func TestBubbleSort(t *testing.T) {
	for _, c := range GetSharedTestCases() {
		original := c.arr
		BubbleSort(c.arr)
		if !slices.Equal(c.arr, c.expected) {
			t.Errorf(
				"BubbleSort(%v) = %v; want %v\n",
				original, c.arr, c.expected,
			)
		}
	}
}

func TestQuickSort(t *testing.T) {
	for _, c := range GetSharedTestCases() {
		original := c.arr
		QuickSort(c.arr)
		if !slices.Equal(c.arr, c.expected) {
			t.Errorf(
				"QuickSort(%v) = %v; want %v\n",
				original, c.arr, c.expected,
			)
		}
	}
}

func TestMergeSort(t *testing.T) {
	for _, c := range GetSharedTestCases() {
		original := c.arr
		MergeSort(c.arr)
		if !slices.Equal(c.arr, c.expected) {
			t.Errorf(
				"MergeSort(%v) = %v; want %v\n",
				original, c.arr, c.expected,
			)
		}
	}
}
