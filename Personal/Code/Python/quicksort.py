# quicksort
def partition(arr,low,high): 
    i = ( low-1 )         # index of smaller element 
    pivot = arr[high]     # pivot 
  
    for j in range(low , high): 
  
        # If current element is smaller than or 
        # equal to pivot 
        if   arr[j] <= pivot: 
          
            # increment index of smaller element 
            i = i+1 
            arr[i],arr[j] = arr[j],arr[i] 
  
    arr[i+1], arr[high] = arr[high], arr[i+1] 
    return ( i+1 ) 
  
def quickSort(arr,low,high): 
    if low < high: 
  
        # pi is partitioning index, arr[p] is now 
        # at right place 
        pi = partition(arr,low,high) 
  
        # Separately sort elements before 
        # partition and after partition 
        quickSort(arr, low, pi-1) 
        quickSort(arr, pi+1, high) 
    else:
        print("Error encountered")

        

def printList(arr): 
    for i in range(len(arr)):         
        print(arr[i],end=" ") 
    print() 
  
  
# Driver code to test above 
def main():
    user_arr = list(input("Give me some numbers seperated by a space\n"))
    print("Given Array: ", end="\n")
    n = len(user_arr) 
    quickSort(user_arr,0,n-1) 
    print ("Sorted array is: ", end="\n")
    printList(user_arr)
    
#     for i in range(n): 
#         sort_list = list([(quickSort(arr,0,n-1))])
#         print ("%d" %arr[i])
main()