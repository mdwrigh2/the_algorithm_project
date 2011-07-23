package sorta

import "testing"

var sortTests = [][]int{
	[]int{5, 2, 4, 6, 1, 3},
	[]int{0, -2, 4, 6, 1, 3},
	[]int{0, 0, 1, 0, 1, 0},
	[]int{5, 4, 3, 2, 1, 0},
	[]int{5, -2, -4, -6, -1, -3},
}

func TestInsertionSort(t *testing.T) {
	for _, st := range sortTests {
		A := make([]int, len(st))
		copy(A, st)
		InsertionSort(A)
		if !IsSorted(A) {
			t.Errorf("sorted %v", st)
			t.Errorf("   got %v", A)
		}
	}
}

func TestMergeSort(t *testing.T) {
	for _, st := range sortTests {
		A := make([]int, len(st))
		copy(A, st)
		MergeSort(A,0,len(A))
		if !IsSorted(A) {
			t.Errorf("sorted %v", st)
			t.Errorf("   got %v", A)
		}
	}
}

func TestBubbleSort(t *testing.T) {
	for _, st := range sortTests {
		A := make([]int, len(st))
		copy(A, st)
		BubbleSort(A)
		if !IsSorted(A) {
			t.Errorf("sorted %v", st)
			t.Errorf("   got %v", A)
		}
	}
}
