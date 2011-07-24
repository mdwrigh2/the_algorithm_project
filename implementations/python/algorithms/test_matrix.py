import unittest
from matrix import matrix_multiply_naive, MatrixException

class MatrixTests(unittest.TestCase):
    def test_matrix_multiply_naive(self):
        mat_a = [[1,2,3], [4,5,6]]
        mat_b = [[1,2],[3,4],[5,6]]
        result = [[22,28], [49,64]] 
        self.assertEquals(result, matrix_multiply_naive(mat_a, mat_b))
        mat_a = [[1,2],[3,4],[5,6]]
        mat_b = [[1,2,3], [4,5,6]]
        result = [[9, 12, 15], [19, 26, 33], [29, 40, 51]]
        self.assertEquals(result, matrix_multiply_naive(mat_a, mat_b))

    def test_matrix_multiply_naive_failure(self):
        mat_a = [[1,2,3,4], [4,5,6,7]]
        mat_b = [[1,2],[3,4],[5,6]]
        self.assertRaises(MatrixException, matrix_multiply_naive, mat_a, mat_b)
