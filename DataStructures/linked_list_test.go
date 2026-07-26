package data_structures

import "testing"

func TestLinkedList_PushAndPop(t *testing.T) {
	tests := []struct {
		name string
		inputs []int
		expected []int
	}{
		{
			name: "Push multiple elements",
			inputs: []int{10, 20, 30},
			expected: []int{30, 20, 10},
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
			list := &LinkedList[int]{}
			for _, v := range tc.inputs {
				list.Push(v)
			}

			if list.Length != len(tc.inputs) {
				t.Errorf("Expected %d, got %d", len(tc.inputs), list.Length)
			}

			for _, expected_val := range tc.expected {
				val, ok := list.Pop()
				if !ok {
					t.Errorf("Expected value %d but list was empty", expected_val)
				}
				if val != expected_val {
					t.Errorf("Expected %d, got %d", expected_val, val)
				}
			}

			_, ok := list.Pop()
			if ok {
				t.Errorf("Expected list to be empty but, pop succeeded")
			}
			if list.Length != 0 {
				t.Errorf("Expected length = 0 but got %d", list.Length)
			}
		})
	}

}
