# mergesort
def mergeSort(arr): 
    if len(arr) >1: 
        mid = len(arr)//2 #Finding the mid of the array 
        L = arr[:mid] # Dividing the array elements  
        R = arr[mid:] # into 2 halves 

        #breakpoint
        mergeSort(L) # Sorting the first half 
        mergeSort(R) # Sorting the second half 
  
        i = j = k = 0
          
       #breakpoint
        while i < len(L) and j < len(R): 
            if L[i] < R[j]: 
                arr[k] = L[i] 
                i+=1
            else: 
                arr[k] = R[j] 
                j+=1
            k+=1
          
        #breakpoint 
        while i < len(L): 
            arr[k] = L[i] 
            i+=1
            k+=1

        #breakpoint  
        while j < len(R): 
            arr[k] = R[j] 
            j+=1
            k+=1
  

def printList(arr): 
    for i in range(len(arr)):         
        print(arr[i],end=" ") 
    print() 
  
# driver code to test the above code 
def main():
    arr = list(input("Give me some numbers seperated by a space\n"))
    print ("Given array is", end="\n")  
    printList(arr) 
    mergeSort(arr) 
    print("Sorted array is:", end="\n") 
    printList(arr) 
main()

