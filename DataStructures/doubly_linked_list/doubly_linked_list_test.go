package doubly_linked_list

import "testing"

func assertPointers[T comparable ](
	t *testing.T,
	l *DoublyLinkedList[T],
	expected []T,
) {
	t.Helper()
	if l.Length != len(expected) {
		t.Fatalf("length mismatch: got %d want %d", l.Length, len(expected))
	}
	if len(expected) == 0 {
		if l.Head != nil || l.Tail != nil {
			t.Errorf("empty list must have nil Head and Tail")
		}
		return
	}
	curr := l.Head
	for i, want := range expected {
		if curr == nil {
			t.Fatalf("forward traversal ended early at index %d", i)
		}
		if curr.Value != want {
			t.Fatalf("forward index %d: got %v want %v", i, curr.Value, want)
		}
		curr = curr.Next
	}
	if curr != nil {
		t.Fatalf("Forward traversal found extra items after the expected tail")
	}

	curr = l.Tail
	for i := len(expected) - 1; i >= 0; i-- {
		if curr == nil {
			t.Fatalf("backward traversal ended early at index %d", i)
		}
		if curr.Value != expected[i] {
			t.Fatalf(
				"backward index %d: got %v want %v",
				i,
				curr.Value,
				expected[i],
			)
		}
		curr = curr.Prev
	}
	if curr != nil {
		t.Fatalf("backward traversal found extra items after the expected head")
	}
}

func TestAppend(t *testing.T) {
	t.Run("appends to an empty list", func(t *testing.T) {
		l := New[int]()
		l.Append(10)
		assertPointers(t, l, []int{10})
	})

	t.Run("appends in correct order", func(t *testing.T) {
		l := New[int]()
		l.Append(10)
		l.Append(30)
		l.Append(60)
		assertPointers(t, l, []int{10, 30, 60})
	})

}

func TestPrepend(t *testing.T) {
	t.Run("prepends to an empty list", func(t *testing.T) {
		l := New[int]()
		l.Prepend(10)
		assertPointers(t, l, []int{10})
	})

	t.Run("prepends in correct order", func(t *testing.T) {
		l := New[int]()
		l.Prepend(10)
		l.Prepend(30)
		l.Prepend(60)
		assertPointers(t, l, []int{60, 30, 10})
	})
}

func TestGet(t *testing.T) {
	l := New[int]()
	l.Append(10)
	l.Append(30)
	l.Append(60)

	tests := []struct {
		name string
		index int
		wantVal int
		wantErr bool
	}{
		{
			name: "get head",
			index: 0,
			wantVal: 10,
			wantErr: false,
		},
		{
			name: "get middle element",
			index: 1,
			wantVal: 30,
			wantErr: false,
		},
		{
			name: "get tail",
			index: 2,
			wantVal: 60,
			wantErr: false,
		},
		{
			name: "get negative index -1",
			index: -1,
			wantVal: 0,
			wantErr: true,
		},
		{
			name: "get index size of length (off by one)",
			index: 3,
			wantVal: 0,
			wantErr: true,
		},
		{
			name: "get large out of bounds",
			index: 100,
			wantVal: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, err := l.Get(tt.index)
			if (err != nil) != tt.wantErr {
				t.Fatalf(
					"Get(%d), error = %v, want = %v", 
					tt.index, 
					err,
					tt.wantErr,
				)
			}
			if !tt.wantErr && val != tt.wantVal {
				t.Errorf("Get(%d) = %v, want %v", tt.index, val, tt.wantVal)
			}
		})
		assertPointers(t, l, []int{10, 30, 60})
	}

}

func TestInsertAt(t *testing.T) {
	l := New[int]()
	l.Append(10)
	l.Append(30)
	l.Append(60)
	l.InsertAt(99, 1)

	assertPointers(t, l, []int{10, 99, 30, 60})
}

func TestRemove(t *testing.T) {
	l := New[int]()
	l.Append(10)
	l.Append(30)
	l.Append(60)
	l.Remove(1)
	assertPointers(t, l, []int{10, 60})
}


