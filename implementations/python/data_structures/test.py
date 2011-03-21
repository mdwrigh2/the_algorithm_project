import unittest
from heaps import MinHeap, MaxHeap

class MinHeapTest(unittest.TestCase):

    def testInitialize(self):
        heap = MinHeap([2,4,6,8,10,1,3,5,7,9])
        self.assertEquals(heap.get(0), 1)

    def testHeapInvariant(self):
        heap = MinHeap([2,4,6,8,10,1,3,5,7,9])
        for k in range(len(heap)):
            if 2*k+1 < len(heap):
                self.assertTrue(heap.get(k) <= heap.get(2*k+1))
            if 2*k+2 < len(heap):
                self.assertTrue(heap.get(k) <= heap.get(2*k+2))


class MaxHeapTest(unittest.TestCase):

    def testInitialize(self):
        heap = MaxHeap([2,4,6,8,10,1,3,5,7,9])
        self.assertEquals(heap.get(0), 10)

    def testHeapInvariant(self):
        heap = MaxHeap([2,4,6,8,10,1,3,5,7,9])
        for k in range(len(heap)):
            if 2*k+1 < len(heap):
                self.assertTrue(heap.get(k) >= heap.get(2*k+1))
            if 2*k+2 < len(heap):
                self.assertTrue(heap.get(k) >= heap.get(2*k+2))

if __name__ == "__main__":
    unittest.main()
