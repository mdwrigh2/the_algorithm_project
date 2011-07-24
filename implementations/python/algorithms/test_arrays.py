import unittest
from arrays import maximum_subarray

class ArrayTests(unittest.TestCase):
    def test_maximum_subarray(self):
        pos_array = [1, 2, 3, 4]
        self.assertEquals((pos_array, sum(pos_array)), maximum_subarray(pos_array))
        neg_array = [-1,-2,-3,-4]
        combined_array = pos_array + neg_array
        self.assertEquals((pos_array, sum(pos_array)),
                          maximum_subarray(combined_array))
        self.assertEquals(([-1], -1), maximum_subarray(neg_array))
        mixed_array = [13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22,
                       15, -4, 7]
        self.assertEquals(([18, 20, -7, 12], 43), maximum_subarray(mixed_array))
