// Alex Ray (2011) <ajray@ncsu.edu>
// TODO(ajray):
// * Support Unicode runes, not just ASCII chars
// * Check sorted while reading in
// * Sort unsorted input
// * Error checking on negative score cases
package sorta

import "testing"

var searchTests = [][]int{
	[]int{1, 2, 3, 4, 5, 6},
	[]int{-8, -2, 0, 6, 11, 13},
}

func TestBinarySearch(t *testing.T) {
	for _, st := range searchTests {
		A := make([]int, len(st))
		copy(A, st) // incase a bad search edits the list
		for i, a := range st {
			if j := BinarySearch(A, a); j != i {
				t.Errorf("searched %d(%d) %v", i,a, st)
				t.Errorf("     got %d(%d) %v", j,j, A)
			}
		}
	}
}
