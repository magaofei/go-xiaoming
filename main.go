package main

import "fmt"

func main() {
	print("HelloWorld")

	arr1 := []int{1,2,3,4}
	fmt.Println(arr1)
	fmt.Println(append(arr1, 5))
	arr2 := make([]int, 4)
	arr3 := make([]int, 4, 40000000000)
	fmt.Println(arr2, cap(arr2), len(arr2))
	fmt.Println(arr3, cap(arr3), len(arr3))
}
