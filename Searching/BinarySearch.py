def BinarySearch(arr,traget):
    low,high = 0,len(arr)-1
    while (low <= high):
        mid = low + (high-low)//2

        if arr[mid] == traget:
            return mid
        elif traget > arr[mid]:
            low = mid +1
        else:
            high = mid -1
    return -1



arr_len = int(input("Enter len of array: "))
arr= []

for i in range (0,arr_len):
    element = int(input(f"Enter element {i+1}: "))
    arr.append(element)

print(arr)

result = BinarySearch(arr,55)

if result == -1:
    print("Element not found")
else:
    print(f"fount in {result}")