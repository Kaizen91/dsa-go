package heap

import  "testing"

func TestPriorityQueue(t *testing.T) {
	tests := []struct {
		name string
		inputs []int
		expected []int
	}{
		{
			name: "Push multiple elements",
			inputs: []int{30, 20, 10},
			expected: []int{10, 20, 30},
		},
		{
			name: "Single element Push and Pop",
			inputs: []int{42},
			expected: []int{42},
		},
		{
			name: "Empty list",
			inputs: []int{},
			expected: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pq := make(PriorityQueue[int], 0)
			pq.Init()
			for _, v := range tc.inputs {
				pq.PushItem(v)
			}

			if pq.Len() != len(tc.inputs) {
				t.Errorf("Expected %d, got %d", len(tc.inputs), pq.Len())
			}

			for _, expected_val := range tc.expected {
				val, ok := pq.PopItem()
				if !ok {
					t.Errorf("Expected value %d but pq was empty", expected_val)
				}
				if val != expected_val {
					t.Errorf("Expected %d, got %d", expected_val, val)
				}
			}

			_, ok := pq.PopItem()
			if ok {
				t.Errorf("Expected pq to be empty but, pop succeeded")
			}
			if pq.Len() != 0 {
				t.Errorf("Expected length = 0 but got %d", pq.Len())
			}

		})
	}
}
