import unittest
from heaps import MinHeap, MaxHeap

class MinHeapTest(unittest.TestCase):

    def test_initialize(self):
        heap = MinHeap([2,4,6,8,10,1,3,5,7,9])
        self.assertEquals(heap.get(0), 1)

    def test_heap_invariant(self):
        heap = MinHeap([2,4,6,8,10,1,3,5,7,9])
        for k in range(len(heap)):
            if 2*k+1 < len(heap):
                self.assertTrue(heap.get(k) <= heap.get(2*k+1))
            if 2*k+2 < len(heap):
                self.assertTrue(heap.get(k) <= heap.get(2*k+2))


class MaxHeapTest(unittest.TestCase):

    def test_initialize(self):
        heap = MaxHeap([2,4,6,8,10,1,3,5,7,9])
        self.assertEquals(heap.get(0), 10)

    def test_heap_invariant(self):
        heap = MaxHeap([2,4,6,8,10,1,3,5,7,9])
        for k in range(len(heap)):
            if 2*k+1 < len(heap):
                self.assertTrue(heap.get(k) >= heap.get(2*k+1))
            if 2*k+2 < len(heap):
                self.assertTrue(heap.get(k) >= heap.get(2*k+2))

if __name__ == "__main__":
    unittest.main()
