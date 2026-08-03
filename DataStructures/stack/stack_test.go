package stack

import "testing"

func TestStack_PushAndPop(t *testing.T) {
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
			name: "Empty stack",
			inputs: []int{},
			expected: []int{},
		},
	}
	
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stack := &Stack[int]{}
			for _, v := range tc.inputs {
				stack.Push(v)
			}

			if stack.Length != len(tc.inputs) {
				t.Errorf("Expected %d, got %d", len(tc.inputs), stack.Length)
			}

			for _, expected_val := range tc.expected {
				val, ok := stack.Pop()
				if !ok {
					t.Errorf("Expected value %d but stack was empty", expected_val)
				}
				if val != expected_val {
					t.Errorf("Expected %d, got %d", expected_val, val)
				}
			}

			_, ok := stack.Pop()
			if ok {
				t.Errorf("Expected stack to be empty but, pop succeeded")
			}
			if stack.Length != 0 {
				t.Errorf("Expected length = 0 but got %d", stack.Length)
			}
		})
	}

}
