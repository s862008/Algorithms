package main

import "fmt"

func main() {
  
  arr := []int{45,55,23,4,56,90,22}
  
  fmt.Println("исходный массив:",arr)
  
  n:=len(arr)
  for i:=0;i<n-1;i++{
	  for j:=0;j<n-i-1;j++{
		  if arr[j]>arr[j+1]{
			  arr[j],arr[j+1] = arr[j+1],arr[j]
		  }
	  }
  }
  
fmt.Println("отсортированный массив:",arr)
  
}
