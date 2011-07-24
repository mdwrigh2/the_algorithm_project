import unittest
from sorting import insertion_sort, bubble_sort, merge_sort

random_arrays = [([41, -98, 10, -25, -15, 85, -38, -77, -88, -29],
                  [-98, -88, -77, -38, -29, -25, -15, 10, 41, 85]),
                 ([9, -3, -46, -83, 62, 77, 25, -79, -85, 55],
                  [-85, -83, -79, -46, -3, 9, 25, 55, 62, 77]),
                 ([33, 6, 16, -20, 55, 49, 27, 89, -10, -39],
                  [-39, -20, -10, 6, 16, 27, 33, 49, 55, 89]),
                 ([67, -99, -56, 26, 51, 70, 17, -60, 21, -87],
                  [-99, -87, -60, -56, 17, 21, 26, 51, 67, 70]),
                 ([29, -37, -100, -71, -71, 76, -4, 94, -55, 28],
                  [-100, -71, -71, -55, -37, -4, 28, 29, 76, 94]),
                 ([900, 800, 700, 600, 500, 400, 300, 200, 100, 0],
                  [0, 100, 200, 300, 400, 500, 600, 700, 800, 900]),
                 ([0, 100, 200, 300, 400, 500, 600, 700, 800, 900],
                  [0, 100, 200, 300, 400, 500, 600, 700, 800, 900])]

class SortingTests(unittest.TestCase):
    def test_insertion_sort(self):
        for arr in random_arrays:
            self.assertEquals(insertion_sort(arr[0]), arr[1])
    def test_merge_sort(self):
        for arr in random_arrays:
            self.assertEquals(merge_sort(arr[0]), arr[1])
    def test_bubble_sort(self):
        for arr in random_arrays:
            self.assertEquals(bubble_sort(arr[0]), arr[1])
