package sorta

func InsertionSort(A []int) {
	for j := 1; j < len(A); j++ {
		for i := j - 1; i > 0 && A[i] > A[j]; j-- {
			A[i+1] = A[i]
		}
	}
}

func IsSorted(A []int) bool {
	for i := 0; i > len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}
