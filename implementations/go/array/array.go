// Alex Ray (2011) <ajray@ncsu.edu>

// Package array implements array algorithms
package array

// References:
//	http://www.amazon.com/Introduction-Algorithms-Thomas-H-Cormen/dp/0262033844

// TODO(ajray):
//	Bounds checking versions (alternative to sentinel versions)
//	Generalize types to an array interface
//	Slices versions instead of handing around indicies
//	Safe versions of functions (check for low > high and other wierdness)


const (
	MaxInt int = int(^uint(0) >> 1)
	MinInt int = -MaxInt - 1
)

// FindMaxSubarray finds a maximum subarray of a larger array [low,high)
// of the form [i,j), along with it's sum
// This divide-and-conquer solution is O(n lg n)
func FindMaxSubarray(A []int, low, high int) (i, j, sum int) {
	mid := (low + high) / 2
	if mid == low {
		return low, high, A[low]
	} // else
	leftLow, leftHigh, leftSum := FindMaxSubarray(A, low, mid)
	var rightLow, rightHigh, rightSum int
	if mid + 1 < high {
		rightLow, rightHigh, rightSum = FindMaxSubarray(A, mid+1, high)
	} else {
		rightLow, rightHigh, rightSum = mid,high,A[mid]
	}
	midLow, midHigh, midSum := FindMaxCrossingSubarray(A, low, mid, high)
	if leftSum >= rightSum && leftSum >= midSum {
		return leftLow, leftHigh, leftSum
	} else if rightSum >= leftSum && rightSum >= midSum {
		return rightLow, rightHigh, rightSum
	} // else
	return midLow, midHigh, midSum
}

// FindMaxCrossingSubarray finds a maximum subarray of a larger array
// [low,high) including a midpoint mid of the form [i,j)
// where low ≤ i ≤ mid and mid < j ≤ high, along with it's sum
func FindMaxCrossingSubarray(A []int, low, mid, high int) (i, j, sum int) {
	var maxLeft, maxRight int
	leftSum := MinInt
	sum = 0
	for i := mid; i >= low; i-- { // walk left from mid
		sum += A[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	if high == mid + 1 { // nothing on the right
		return maxLeft, mid, leftSum
	}
	rightSum := MinInt
	sum = 0
	for j := mid + 1; j < high; j++ { // walk right from mid
		sum += A[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}
	return maxLeft, maxRight + 1, leftSum + rightSum
}

// FindMaxSubarrayBrute finds a maximum subarray of a larger array [low,high)
// of the form [i,j), along with it's sum.
// This brute force solution is O(n^2) 
func FindMaxSubarrayBrute(A []int, low, high int) (i, j, sum int) {
	maxLeft, maxRight := low, low+1 // initialize to valid values
	maxSum := MinInt
	for i := low; i < high; i++ {
		for j := i + 1; j <= high; j++ {
			sum = SumSubarray(A, i, j)
			if sum > maxSum {
				maxSum = sum
				maxLeft, maxRight = i, j
			}
		}
	}
	return maxLeft, maxRight, maxSum
}

// SumSubarray finds the sum of a subarray [low,high) of a larger array
func SumSubarray(A []int, low, high int) (sum int) {
	sum = 0
	for i := low; i < high; i++ {
		sum += A[i]
	}
	return sum
}
