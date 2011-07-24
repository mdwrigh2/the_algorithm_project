class MatrixException(Exception):
    pass

# RUNTIME: O(n^3)
def matrix_multiply_naive(mat_a, mat_b):
    if len(mat_a[0]) != len(mat_b):
        raise MatrixException("Unable to multiply matrices: improper sizes")
    c = []
    for i in range(len(mat_a)):
        c.append([])
        for j in range(len(mat_b[0])):
            c[i].append(0)
            for k in range(len(mat_b)):
                c[i][j] += mat_a[i][k] * mat_b[k][j]
    return c
