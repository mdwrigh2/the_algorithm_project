# Essentially this problem deals with finding the subsequence of an array of
# numbers that has the largest sum. This is only interesting when the given
# array contains negative numbers, otherwise the answer will always be the
# entire array.
# This function will return a tuple containing the maximum subarray, and its
# sum.
# RUNTIME: O(n lg(n))
def maximum_subarray(arr):
    if len(arr) == 1:
        return (arr, arr[0])
    mid = len(arr)/2
    lower = maximum_subarray(arr[:mid])
    upper = maximum_subarray(arr[mid:])
    # So that covers if the max sub array exists in either the left or right
    # sections of the array. It could also cross the middle, so we have to find
    # the max subarray that cross the middle section
    left_sum = arr[mid]
    temp_sum = arr[mid]
    max_left = mid
    for i in reversed(range(mid)):
        temp_sum += arr[i]
        if temp_sum > left_sum:
            left_sum = temp_sum
            max_left = i
    right_sum = arr[mid]
    temp_sum = arr[mid]
    max_right = mid
    for i in range(mid + 1, len(arr)):
        temp_sum += arr[i]
        if temp_sum > right_sum:
            right_sum = temp_sum
            max_right = i
    # Add one to max right because array slices are non-inclusive for the
    # upperbound.
    # Also, subtract out arr[mid] because it's in both left and right sums.
    cross_mid = (arr[max_left:max_right + 1], left_sum + right_sum - arr[mid])
    if lower[1] >= upper[1] and lower[1] >= cross_mid[1]:
        return lower
    elif upper[1] >= lower[1] and upper[1] >= cross_mid[1]:
        return upper
    else:
        return cross_mid
