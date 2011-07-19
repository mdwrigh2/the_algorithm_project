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
