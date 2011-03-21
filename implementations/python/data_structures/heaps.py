import heapq

class MinHeap():
    ''' A minimum heap. This is kind of cheating since I use the built-in heap implementation '''
    def __init__(self, items = None):
        if not items:
            self.heap = []
        else:
            self.heap = items[:]
        heapq.heapify(self.heap)

    def push(self, item):
        heapq.heappush(self.heap, item)

    def pop(self):
        return heapq.heappop(self.heap)

    def get(self, num):
        return self.heap[num]

    def size(self):
        return len(self.heap)

    def __len__(self):
        return len(self.heap)

class MaxHeap():
    ''' A maximum heap. This is kind of cheating since I use the built-in heap implementation '''
    def __init__(self, items = None):
        if not items:
            self.heap = []
        else:
            self.heap = map(lambda x :-x, items)
        heapq.heapify(self.heap)

    def push(self, item):
        heapq.heappush(self.heap, -item)

    def pop(self):
        return -heapq.heappop(self.heap)

    def get(self, num):
        return -self.heap[num]

    def size(self):
        return len(self.heap)

    def __len__(self):
        return len(self.heap)
