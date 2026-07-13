package search

import "testing"

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
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
	for _, c := range cases {
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
