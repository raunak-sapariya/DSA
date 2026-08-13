def LinearSearch(arr,len,traget):
    for i in range(0,len):
        if arr[i]==traget:
            return i
    return -1


arr = [5,9,7,3,65,55,94]
len = len(arr)
target = 6588

result = LinearSearch(arr,len,target)

if result == -1:
    print("Element not found")
else:
    print(f"fount in {result}")