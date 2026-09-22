def counting_sort(arr, exp):
    n = len(arr)
    output = [0] * n
    count = [0] * 10

    # count of each digit
    for num in arr:
        digit = (num // exp) % 10
        count[digit] += 1

    # prefix sum
    for i in range(1, 10):
        count[i] += count[i- 1]

    # output from right to left
    for i in range(n-1, -1, -1):
        digit = (arr[i] // exp) % 10
        count[digit] -= 1
        output[count[digit]] = arr[i]

    for i in range(n):
        arr[i] = output[i]

def radix(arr):
    if not arr:
        return
    max = arr[0]
    for i in arr:
        if max < i:
            max = i
    exp =1

    while max // exp > 0:
        counting_sort(arr, exp)
        exp *= 10

if __name__ == "__main__":
    
    arr = [29, 83, 471, 36, 91, 8]
    print("Original array:", arr)
    radix(arr)
    print("Sorted array:", arr)