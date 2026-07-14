package search

import "testing"

type SharedTestCase struct {
		name     string
		arr      []int
		target   int
		expected int
}

func GetSharedCases() []SharedTestCase {
	return []SharedTestCase{
		{
			name:     "target is in arr",
			arr:      []int{1, 2, 3, 4, 5},
			target:   3,
			expected: 2,
		},
		{
			name:     "target is at end",
			arr:      []int{1, 2, 3, 4, 9},
			target:   9,
			expected: 4,
		},
		{
			name:     "target is at begining",
			arr:      []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "target not in array",
			arr:      []int{1, 2, 3, 4, 5},
			target:   6,
			expected: -1,
		},
		{
			name:     "empty array",
			arr:      []int{},
			target:   6,
			expected: -1,
		},
	}
}

func TestLinearSearch(t *testing.T) {
	for _, c := range GetSharedCases() {
		t.Run(c.name, func(t *testing.T) {
			actual := LinearSearch(c.target, c.arr)
			if actual != c.expected {
				t.Errorf(
					"LinearSearch(%d, %v) = %d want %d",
					c.target,
					c.arr,
					actual,
					c.expected,
				)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	for _, c := range GetSharedCases() {
		t.Run(c.name, func(t *testing.T) {
			actual := BinarySearch(c.target, c.arr)
			if actual != c.expected {
				t.Errorf(
					"BinarySearch(%d, %v) = %d want %d",
					c.target,
					c.arr,
					actual,
					c.expected,
				)
			}
		})
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	for _, c := range GetSharedCases() {
		t.Run(c.name, func(t *testing.T) {
			actual := BinarySearchRecursive(c.target, c.arr)
			if actual != c.expected {
				t.Errorf(
					"BinarySearchRecursive(%d, %v) = %d want %d",
					c.target,
					c.arr,
					actual,
					c.expected,
				)
			}
		})
	}
}
