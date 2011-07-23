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
        if len(left) == i or left[i] > right[j]:
            sorted_array.append(right[j])
            j += 1
        else:
            sorted_array.append(left[i])
            i += 1
