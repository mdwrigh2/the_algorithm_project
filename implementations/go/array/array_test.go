// Alex Ray (2011) <ajray@ncsu.edu>

package array

// TODO(ajray):
//	Randomized testing D&C vs Brute algorithms
//	Benchmarking

import "testing"

type arraySum struct {
	array       []int
	left, right int
	sum         int
}

var arrayTests = []arraySum{
	arraySum{[]int{-5, -2, -4, -6, -1, -3}, 4, 5, -1},
	arraySum{[]int{5, 2, 4, 6, 1, 3}, 0, 6, 21},
	arraySum{[]int{0, -2, 4, 6, 1, 3}, 2, 6, 14},
	arraySum{[]int{5, -4, 3, -2, 1, 0}, 0, 1, 5},
}

func TestFindMaxSubarray(t *testing.T) {
	for _, A := range arrayTests {
		left, right, sum := FindMaxSubarray(A.array,0,len(A.array))
		if A.left != left || A.right != right || A.sum != sum{
			t.Errorf("FindMaxSubarray got [%d,%d):%d", left,right,sum)
			t.Errorf("    expected [%d,%d):%d", A.left,A.right,A.sum)
			t.Errorf("    array %v", A.array)
		}
	}
}

func TestFindMaxSubarrayBrute(t *testing.T) {
	for _, A := range arrayTests {
		left, right, sum := FindMaxSubarrayBrute(A.array,0,len(A.array))
		if A.left != left || A.right != right || A.sum != sum{
			t.Errorf("FindMaxSubarrayBrute got [%d,%d):%d", left,right,sum)
			t.Errorf("    expected [%d,%d):%d", A.left,A.right,A.sum)
			t.Errorf("    array %v", A.array)
		}
	}
}
