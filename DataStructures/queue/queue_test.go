package queue

import "testing"

func TestQueue_PushAndPop(t *testing.T) {
	tests := []struct {
		name string
		inputs []int
		expected []int
	}{
		{
			name: "Push multiple elements",
			inputs: []int{10, 20, 30},
			expected: []int{10, 20, 30},
		},
		{
			name: "Single element Push and Pop",
			inputs: []int{42},
			expected: []int{42},
		},
		{
			name: "Empty queue",
			inputs: []int{},
			expected: []int{},
		},
	}
	
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			queue := &Queue[int]{}
			for _, v := range tc.inputs {
				queue.Push(v)
			}

			if queue.Length != len(tc.inputs) {
				t.Errorf("Expected %d, got %d", len(tc.inputs), queue.Length)
			}

			for _, expected_val := range tc.expected {
				val, ok := queue.Pop()
				if !ok {
					t.Errorf("Expected value %d but queue was empty", expected_val)
				}
				if val != expected_val {
					t.Errorf("Expected %d, got %d", expected_val, val)
				}
			}

			_, ok := queue.Pop()
			if ok {
				t.Errorf("Expected queue to be empty but, pop succeeded")
			}
			if queue.Length != 0 {
				t.Errorf("Expected length = 0 but got %d", queue.Length)
			}
		})
	}

}
