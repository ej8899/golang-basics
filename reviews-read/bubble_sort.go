package main

import "fmt"

func Swap(array []int, idx int) {
    array[idx], array[idx+1] = array[idx+1], array[idx]
}

func BubbleSort(array []int) {
    for _, _ = range(array) {
        for idx := 0; idx + 1 < len(array); idx ++ {
            if array[idx] > array[idx+1] {
                Swap(array, idx)
            }
        }
    }
}

func main() {
  fmt.Println("Enter an array length:")
  var arrLen int
  _, _ = fmt.Scanf("%d", &arrLen)
  arr := make([]int, arrLen, arrLen)
  fmt.Println("Enter array elements (every element in a separate line):")
  for i := 0; i < arrLen; i ++ {
     _, _ = fmt.Scanf("%d", &arr[i])
  }

  BubbleSort(arr)
  fmt.Println("Sorted array:")

  for _, el := range(arr) {
      fmt.Printf("%d ", el)
  }
  fmt.Println()
}