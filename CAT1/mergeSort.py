def merge(arr, low, mid, high):
    # print("Merging subarrays:", arr[low:mid + 1], "and", arr[mid + 1:high + 1])
    # print("Low:", low, "Mid:", mid, "High:", high)
    temp = []
    left = low
    right = mid + 1

    while left <= mid and right <= high:
        if arr[left] <= arr[right]:
            temp.append(arr[left])
            left += 1
        else:
            temp.append(arr[right])
            right += 1

    # Add remaining elements from the left subarray
    while left <= mid:
        temp.append(arr[left])
        left += 1

    # Add remaining elements from the right subarray
    while right <= high:
        temp.append(arr[right])
        right += 1

    # Copy the sorted elements back to the original array
    # print("Merged array:", temp)
    for i in range(len(temp)):
        arr[low + i] = temp[i]
    # print("Array after merging:", arr[low:high + 1])
    # print()

def mergesSort(arr, low, high):
    if low >= high:
        return
    mid = (low + high) // 2
    mergesSort(arr, low, mid)
    mergesSort(arr, mid + 1, high)
    merge(arr, low, mid, high)

arr = [1,6,7,4,2,9,8,5,3]
mergesSort(arr, 0, len(arr) - 1)
print("Sorted array is:", arr)