package main

import (
	"fmt"
	"technical-test/problem"
)

func main() {
	arr1 := []int{1, 2, 3, 8, 9, 3, 2, 1} // result -> 3
	arr2 := []int{1, 2, 1, 2, 2, 1}       // result -> 2
	arr3 := []int{7, 1, 2, 9, 7, 2, 1}    // result -> 2
	fmt.Println(problem.MaxNumberArray(arr1))
	fmt.Println(problem.MaxNumberArray(arr2))
	fmt.Println(problem.MaxNumberArray(arr3))
}
