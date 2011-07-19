package sorta

import "testing"

var aTest = [][]int{
	[]int{5, 2, 4, 6, 1, 3},
	[]int{0,-2, 4, 6, 1, 3},
	[]int{0, 0, 1, 0, 1, 0},
	[]int{5, 4, 3, 2, 1, 0},
	[]int{5,-2,-4,-6,-1,-3},
}

func TestInsertionSort(t *testing.T) {
	for _, A := range aTest {
		InsertionSort(A)
		if !IsSorted(A) {
			t.Errorf("InsertionSort returned %v", A)
		}
	}
}
