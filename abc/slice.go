package main

import "fmt"

func main() {
	var slice []int
	var slice1 []int = make([]int, 2)

	slice2 := make([]int, 3)

	slice3 := []int{1, 2, 3}
	slice4 := slice3[:]

	fmt.Printf("%d, %d, %d, %d, %d\n, %v", len(slice), len(slice1), len(slice2), len(slice3), len(slice4), slice4)
}
