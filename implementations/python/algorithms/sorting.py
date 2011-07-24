# Insertion sort is a very naive sorting algorithm. Essentially it creates a
# sorted sublist by traversing the list and inserting each element into the
# sublist that is all elements previous to itself, in a sorted fashion. This
# means that no matter what element you're on, all elements previous to it are
# sorted. The insertion of the current element in this implementation is just a
# linear search (i.e. traverse through the sorted sublist until you find the
# proper location to insert the item). You could use binary search to find the
# insertion point (which is known as binary insertion sort), but because you
# have to shift elements in the array in order to insert the current element in
# its proper location, the runtime doesn't change. Similarly, using a Linked
# List allows for fast insertion, but slow search. Using more advanced data
# structures, such as heaps, can improve the running time of this algorithm, and
# is the basis for sorting methods such as heap sort.
# RUNTIME: O(n^2), where n = len(arr)
def insertion_sort(arr):
    # Copy the array so we don't modify the original array
    sorted_arr = arr[:]
    for i in range(1, len(arr)):
        key = sorted_arr[i]
        j = i - 1
        while j >= 0 and sorted_arr[j] > key:
            # Shift the element to the position to its right
            sorted_arr[j+1] = sorted_arr[j]
            # Move to the next element
            j -= 1
        sorted_arr[j+1] = key
    return sorted_arr

def merge_sort(arr):
    size = len(arr)
    # Base case!
    if size == 1:
        return arr
    # Divide!
    left = merge_sort(arr[:size/2])
    right = merge_sort(arr[size/2:])
    # Merge! (aka. Conquer!)
    sorted_array = []
    i = 0
    j = 0
    for k in range(len(left) + len(right)):
        # Make sure we're still within the bounds of the arrays
        if len(left) == i:
            sorted_array.append(right[j])
            j += 1
        elif len(right) == j:
            sorted_array.append(left[i])
            i += 1
        # If we are, grab the smaller of the two elements and increment the
        # index for that array to signify we've used that value
        elif left[i] > right[j]:
            sorted_array.append(right[j])
            j += 1
        else:
            sorted_array.append(left[i])
            i += 1
    return sorted_array

def bubble_sort(arr):
    sorted_arr = arr[:]
    for i in range(len(sorted_arr)-1):
        for j in reversed(range(i+1, len(sorted_arr))):
            if sorted_arr[j] < sorted_arr[j-1]:
                sorted_arr[j], sorted_arr[j-1] = sorted_arr[j-1], sorted_arr[j]
    return sorted_arr
