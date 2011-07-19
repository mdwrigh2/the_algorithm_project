package sorta

const MaxInt = int(^uint(0) >> 1)

func InsertionSort(A []int) {
	for j := 1; j < len(A); j++ {
		key := A[j]
		var i int
		for i = j - 1; i >= 0 && A[i] > key; i-- {
			A[i+1] = A[i]
		}
		A[i+1] = key
	}
}

func IsSorted(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}
