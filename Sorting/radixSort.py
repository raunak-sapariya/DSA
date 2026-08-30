def count_sort(arr, n, pos):
    count = [0] * 10
    output = [0] * n

    for i in range(0, n):
        count[(arr[i] // pos) % 10] += 1

    for i in range(1, 10):
        count[i] += count[i - 1]

    i = n - 1
    while i >= 0:
        output[count[(arr[i] // pos) % 10] - 1] = arr[i]
        count[(arr[i] // pos) % 10] -= 1
        i -= 1

    for i in range(0, n):
        arr[i] = output[i]

def radix_sort(arr):
    max1 = max(arr)
    pos = 1
    while max1 // pos > 0:
        count_sort(arr, len(arr), pos)
        pos *= 10

if __name__ == "__main__":
    arr = [201, 175, 100, 142, 158, 165, 112, 190, 105]
    print("Original array:", arr)
    radix_sort(arr)
    print("Sorted array:", arr)