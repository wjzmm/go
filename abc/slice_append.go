package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	numbers = append(numbers, 5)
	numbers = append(numbers, 6)

	numbers_bigger := make([]int, len(numbers), (cap(numbers) * 2))

	copy(numbers_bigger, numbers)

	fmt.Printf("len = %d, %v", len(numbers_bigger), numbers_bigger)
}
