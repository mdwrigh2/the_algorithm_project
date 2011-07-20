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

// TODO: parallelize this?
// TODO: assumes start < end
// sort items in [start,end)
func MergeSort(A []int, start, end int) {
	half := (start + end) / 2
	if half > start {
		MergeSort(A, start, half)
		MergeSort(A, half, end)
		Merge(A, start, half, end)
	}
}

// TODO: assumes start < half < end
func Merge(A []int, start, half, end int) {
	n1 := half - start          // length of first stack
	n2 := end - half            // length of second stack
	first := make([]int, n1+1)  // extra item for 'infinity'
	second := make([]int, n2+1) //   is the end of the lists
	for i := 0; i < n1; i++ {
		first[i] = A[start+i]
	}
	for j := 0; j < n2; j++ {
		second[j] = A[half+j]
	}
	first[n1] = MaxInt  // marks the end of the stacks
	second[n2] = MaxInt // saves us from having to bounds-check
	for k := 0; k < n1+n2; k++ {
		i := 0
		j := 0
		if first[i] < second[j] {
			A[k] = first[i]
			i++
		} else {
			A[k] = second[j]
			j++
		}
	}
}

func BubbleSort(A []int) {
	for i := 0 ; i < len(A) ; i++ {
		for j := i ; j > 0 ; j-- {
			if A[j-1] > A[j] {
				A[j-1], A[j] = A[j], A[j-1]
			}
		}
	}
}
