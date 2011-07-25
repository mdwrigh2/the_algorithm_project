// Alex Ray (2011) <ajray@ncsu.edu>
// TODO(ajray):
// * Support Unicode runes, not just ASCII chars
// * Check sorted while reading in
// * Sort unsorted input
// * Error checking on negative score cases
package sorta


func BinarySearch(A []int, x int) (i int) {
	min, max := 0, len(A)
	for min <= max {
		mid := (min + max) / 2
		if A[mid] == x {
			return mid
		}
		if x > A[mid] {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1
}

