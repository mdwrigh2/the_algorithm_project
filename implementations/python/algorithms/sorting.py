def insertion_sort(arr):
    sorted_arr = arr[:]
    for i in range(1, len(arr)):
        key = sorted_arr[i]
        j = i - 1
        while j >= 0 and sorted_arr[j] > key:
            sorted_arr[j+1] = sorted_arr[j]
            j = j - 1
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
