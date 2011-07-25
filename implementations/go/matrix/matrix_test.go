// Alex Ray (2011) <ajray@ncsu.edu>

package matrix

// TODO(ajray):
//	Randomized testing
//	Benchmarking

import "testing"
//import "fmt"

var (
	pintMatrix = Matrix([]int{1,0,0,0, 0,1,0,0, 0,0,1,0, 0,0,0,1})
	pentMatrix = Matrix([]int{2,0,0,0, 0,2,0,0, 0,0,2,0, 0,0,0,2})
	pestMatrix = Matrix([]int{8,3,1,9, 9,4,7,5, 6,6,32,87, 21,45,66,99})
	testMatrix = Matrix([]int{0,1,2,3, 4,5,6,7, 8,9,10,11, 12,13,14,15})
	tastMatrix = Matrix([]int{1,1,2,3, 4,6,6,7, 8,9,11,11, 12,13,14,16})
	fastMatrix = Matrix([]int{8,4,3,12, 13,9,13,12, 14,15,42,98, 33,58,80,114})
	pastMatrix = Matrix([]int{84,151,269,476,260,383,693,1276,436,615,1117,2076,612,847,1541,2876})
	partMatrix = []Matrix{
		Matrix([]int{0,1,4,5}),
		Matrix([]int{2,3,6,7}),
		Matrix([]int{8,9,12,13}),
		Matrix([]int{10,11,14,15}),
	}
)

func TestSquareMatrixPartition(t *testing.T) {
	//fmt.Println(testMatrix.String(4,4))
	A,B,C,D,i,j := SquareMatrixPartition(testMatrix, 4)
	//fmt.Println(A.String(i,i))
	//fmt.Println(B.String(i,j))
	//fmt.Println(C.String(j,i))
	//fmt.Println(D.String(j,j))
	if !A.Equal(partMatrix[0],i,j) {
		t.Errorf("SquareMatrixPartition got:\n%v expected:\n%v", 
			A.String(i,j),partMatrix[0].String(i,j))
	}
	if !B.Equal(partMatrix[1],i,j) {
		t.Errorf("SquareMatrixPartition got:\n%v expected:\n%v", 
			B.String(i,j),partMatrix[1].String(i,j))
	}
	if !C.Equal(partMatrix[2],i,j) {
		t.Errorf("SquareMatrixPartition got:\n%v expected:\n%v", 
			C.String(i,j),partMatrix[2].String(i,j))
	}
	if !D.Equal(partMatrix[3],i,j) {
		t.Errorf("SquareMatrixPartition got:\n%v expected:\n%v", 
			D.String(i,j),partMatrix[3].String(i,j))
	}
}

func TestSquareMatrixMerge(t *testing.T) {
	A,n := SquareMatrixMerge(partMatrix[0],partMatrix[1],partMatrix[2],partMatrix[3],2)
	if !A.Equal(testMatrix,n,n) {
		t.Errorf("SquareMatrixMerge got:\n%v expected:\n%v", 
			A.String(n,n),testMatrix.String(n,n))
	}
}

func TestSquareMatrixAdd(t *testing.T) {
	// identity + identity matrix
	C := SquareMatrixAdd(pintMatrix, pintMatrix, 4)
	if !C.Equal(pentMatrix,4,4) {
		t.Errorf("SquareMatrixAdd got:\n%v expected:\n%v", 
			C.String(4,4),pentMatrix.String(4,4))
	}
	// test identity matrix
	C = SquareMatrixAdd(testMatrix, pintMatrix, 4)
	if !C.Equal(tastMatrix,4,4) {
		t.Errorf("SquareMatrixAdd got:\n%v expected:\n%v", 
			C.String(4,4),tastMatrix.String(4,4))
	}
	// test arbitrary matrix
	C = SquareMatrixAdd(testMatrix, pestMatrix, 4)
	if !C.Equal(fastMatrix,4,4) {
		t.Errorf("SquareMatrixAdd got:\n%v expected:\n%v", 
			C.String(4,4),fastMatrix.String(4,4))
	}
}

func TestSquareMatrixMultiply(t *testing.T) {
	// test identity matrix
	C := SquareMatrixMultiply(testMatrix, pintMatrix, 4)
	if !C.Equal(testMatrix,4,4) {
		t.Errorf("SquareMatrixMultiply got:\n%v expected:\n%v", 
			C.String(4,4),testMatrix.String(4,4))
	}
	// test arbitrary matrix
	C = SquareMatrixMultiply(testMatrix, pestMatrix, 4)
	if !C.Equal(pastMatrix,4,4) {
		t.Errorf("SquareMatrixMultiply got:\n%v expected:\n%v", 
			C.String(4,4),pastMatrix.String(4,4))
	}
}

func TestSquareMatrixMultiplyRecursive(t *testing.T) {
	// test identity matrix
	C := SquareMatrixMultiplyRecursive(testMatrix, pintMatrix, 4)
	if !C.Equal(testMatrix,4,4) {
		t.Errorf("SquareMatrixMultiplyRecursive got:\n%v expected:\n%v", 
			C.String(4,4),testMatrix.String(4,4))
	}
	// test arbitrary matrix
	C = SquareMatrixMultiplyRecursive(testMatrix, pestMatrix, 4)
	if !C.Equal(pastMatrix,4,4) {
		t.Errorf("SquareMatrixMultiplyRecursive got:\n%v expected:\n%v", 
			C.String(4,4),pastMatrix.String(4,4))
	}
}
