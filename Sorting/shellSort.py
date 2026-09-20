"""
Time Complexity:
    Best Case: O(n log n), when the array is already nearly sorted (depends on gap sequence).
    Average Case: Between O(n^1.25) and O(n^1.5), depending on the chosen gap sequence.
    Worst Case: O(n²), with simple gap sequences like n/2, n/4, …, 1.
Space Complexity: O(1)
"""
def shellSort(arr):
    n = len(arr)
    gap = n // 2

    while gap > 0:
        for i in range(gap, n):
            temp = arr[i]
            j = i
            while j >= gap and arr[j - gap] > temp:
                arr[j] = arr[j - gap]
                j -= gap
            arr[j] = temp
        gap //= 2

    return arr

arr = [12, 34, 54, 2, 3]
sorted_arr = shellSort(arr)
print("Sorted array:", sorted_arr)