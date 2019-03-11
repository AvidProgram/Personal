
# Bubble sort implementation 
 
def bubbleSort(arr):
    n = len(arr)
    # Traverse through all array elements
    for i in range(n):
 
        # Last i elements are already in place
        for j in range(0, n-i-1, 1):
            if arr[j] > arr[j+1]:
                #swap if true
                arr[j], arr[j+1] = arr[j+1], arr[j] 

arr = [6, 8, 10, 4, 3, 5]
bubbleSort(arr)
print ("Sorted array is:")
for i in range(len(arr)):
    print ("%d" %arr[i])