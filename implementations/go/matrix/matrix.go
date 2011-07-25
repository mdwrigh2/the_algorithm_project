// Alex Ray (2011) <ajray@ncsu.edu>

// Package matrix implements matrix algorithms
package matrix

// References:
//	http://www.amazon.com/Introduction-Algorithms-Thomas-H-Cormen/dp/0262033844

// TODO(ajray):
//	Version with deep ([][]int) instead of flat ([]int) matricies
//	Generalize types to an matrix interface
//	Slices versions instead of handing around indicies
//	Safe versions of functions (check for wierdness)
//	Use indicies instead of copying matricies

import "fmt"

// Matrix is a flat matrix, so access the ith,jth element
// by A[i*n+j] where n is the number of rows
type Matrix []int

// String returns a prettyprinted representation of the m*n Matrix 
func (A Matrix) String(m, n int) (s string) {
	s = ""
	for i := 0 ; i < m ; i++ {
		for j := 0 ; j < n ; j++ {
			s += fmt.Sprintf("%d ",A[i*m+j])
		}
		s += "\n"
	}
	return s
}

// Equal checks of the m*n Matricies are element-wise equal
func (A Matrix) Equal(B Matrix, m, n int) bool {
	for i := 0 ; i < m ; i++ {
		for j := 0 ; j < n ; j++ {
			if A[i*m+j] != B[i*m+j] {
				return false
			}
		}
	}
	return true
}

// SquareMatrixAdd adds two n*n matricies.
// It does NOT check for squareness
func SquareMatrixAdd(A, B Matrix, n int) (C Matrix) {
	C = Matrix(make([]int,n*n))
	for i := 0 ; i < n ; i++ {
		for j := 0 ; j < n ; j++ {
			C[i*n+j] += A[i*n+j] + B[i*n+j]
		}
	}
	return C
}

// SquareMatrixSub subtracts two n*n matricies.
// It does NOT check for squareness
func SquareMatrixSub(A, B Matrix, n int) (C Matrix) {
	C = Matrix(make([]int,n*n))
	for i := 0 ; i < n ; i++ {
		for j := 0 ; j < n ; j++ {
			C[i*n+j] += A[i*n+j] - B[i*n+j]
		}
	}
	return C
}

// SquareMatrixMultiply multiplies two n*n matricies.
// It does NOT check for squareness
func SquareMatrixMultiply(A, B Matrix, n int) (C Matrix) {
	C = Matrix(make([]int,n*n))
	for i := 0 ; i < n ; i++ {
		for j := 0 ; j < n ; j++ {
			for k := 0 ; k < n ; k++ {
				C[i*n+j] += A[i*n+k] * B[k*n+j]
			}
		}
	}
	return C
}

// SquareMatrixPartition partitions a n*n matrix into four
// mutually exclusive and exhaustive submatricies.
// ONLY WORKS FOR n = POWERS OF 2
func SquareMatrixPartition(A Matrix, n int) (A11,A12,A21,A22 Matrix, i, j int) {
	half := n / 2
	A11 = Matrix(make([]int,half*half))
	A12 = Matrix(make([]int,half*(n-half)))
	A21 = Matrix(make([]int,(n-half)*half))
	A22 = Matrix(make([]int,(n-half)*(n-half)))
	for i := 0 ; i < half ; i++ {
		for j := 0 ; j < half ; j++ {
			A11[i*half+j] = A[i*n+j]
		}
		for j := 0 ; j < n-half ; j++ {
			A12[i*half+j] = A[i*n+j+half]
		}
	}
	for i := 0 ; i < n-half ; i++ {
		for j := 0 ; j < half ; j++ {
			A21[i*(n-half)+j] = A[(i+half)*n+j]
		}
		for j := 0 ; j < n-half ; j++ {
			A22[i*(n-half)+j] = A[(i+half)*n+j+half]
		}
	}
	return A11, A12, A21, A22, half, n-half
}

// SquareMatrixMerge merges 4 n*n Matricies into one 2n*2n Matrix
func SquareMatrixMerge(A,B,C,D Matrix, n int) (E Matrix, m int) {
	E = Matrix(make([]int,4*n*n))
	for i := 0 ; i < n ; i++ {
		for j := 0 ; j < n ; j++ {
			E[i*2*n+j] = A[i*n+j]
			E[i*2*n+j+n] = B[i*n+j]
			E[(i+n)*2*n+j] = C[i*n+j]
			E[(i+n)*2*n+j+n] = D[i*n+j]
		}
	}
	return E, 2*n
}

// SquareMatrixMultiplyRecursive mutiplies two n*n matricies
// It does NOT check for squareness
// ONLY WORKS FOR n = POWERS OF 2
func SquareMatrixMultiplyRecursive(A, B Matrix, n int) (C Matrix) {
	C = Matrix(make([]int,n*n))
	if n == 1 {
		C[0] = A[0]*B[0]
	} else {
		A11,A12,A21,A22,_,_ := SquareMatrixPartition(A,n)
		B11,B12,B21,B22,_,_ := SquareMatrixPartition(B,n)
		C11 := SquareMatrixAdd(SquareMatrixMultiplyRecursive(A11,B11,n/2),
			SquareMatrixMultiplyRecursive(A12,B21,n/2),n/2)
		C12 := SquareMatrixAdd(SquareMatrixMultiplyRecursive(A11,B12,n/2),
			SquareMatrixMultiplyRecursive(A12,B22,n/2),n/2)
		C21 := SquareMatrixAdd(SquareMatrixMultiplyRecursive(A21,B11,n/2),
			SquareMatrixMultiplyRecursive(A22,B21,n/2),n/2)
		C22 := SquareMatrixAdd(SquareMatrixMultiplyRecursive(A21,B12,n/2),
			SquareMatrixMultiplyRecursive(A22,B22,n/2),n/2)
		C,_ = SquareMatrixMerge(C11,C12,C21,C22,n/2)
	}
	return C
}

// SquareMatrixMultiplyStrassen mutiplies two n*n matricies
// It does NOT check for squareness
// ONLY WORKS FOR n = POWERS OF 2
func SquareMatrixMultiplyStrassen(A, B Matrix, n int) (C Matrix) {
	C = Matrix(make([]int,n*n))
	if n == 1 {
		C[0] = A[0]*B[0]
	} else {
		A11,A12,A21,A22,_,_ := SquareMatrixPartition(A,n)
		B11,B12,B21,B22,_,_ := SquareMatrixPartition(B,n)
		S1 := SquareMatrixSub(B12,B22,n/2)
		S2 := SquareMatrixAdd(A11,A12,n/2)
		S3 := SquareMatrixAdd(A21,A22,n/2)
		S4 := SquareMatrixSub(B21,B11,n/2)
		S5 := SquareMatrixAdd(A11,A22,n/2)
		S6 := SquareMatrixAdd(B11,B22,n/2)
		S7 := SquareMatrixSub(A12,A22,n/2)
		S8 := SquareMatrixAdd(B21,B22,n/2)
		S9 := SquareMatrixSub(A11,A21,n/2)
		S10:= SquareMatrixAdd(B11,B12,n/2)
		P1 := SquareMatrixMultiplyStrassen(A11,S1, n/2)
		P2 := SquareMatrixMultiplyStrassen(S2, B22,n/2)
		P3 := SquareMatrixMultiplyStrassen(S3, B11,n/2)
		P4 := SquareMatrixMultiplyStrassen(A22,S4, n/2)
		P5 := SquareMatrixMultiplyStrassen(S5, S6, n/2)
		P6 := SquareMatrixMultiplyStrassen(S7, S8, n/2)
		P7 := SquareMatrixMultiplyStrassen(S9, S10,n/2)
		C11 := SquareMatrixAdd(SquareMatrixAdd(P5,P4,n/2),
			SquareMatrixSub(P6,P2,n/2),n/2)
		C12 := SquareMatrixAdd(P1,P2,n/2)
		C21 := SquareMatrixAdd(P3,P4,n/2)
		C22 := SquareMatrixSub(SquareMatrixAdd(P5,P1,n/2),
			SquareMatrixAdd(P3,P7,n/2),n/2)
		C,_ = SquareMatrixMerge(C11,C12,C21,C22,n/2)
	}
	return C
}
