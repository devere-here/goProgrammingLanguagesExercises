package main

import "fmt"

func main() {
	arr := [6]int{12, 3, 42, 2, 6, 19}
	slice := []int{12, 3, 42, 2, 6, 19}

	reverseArray(&arr)
	fmt.Println("reversing the array", arr)

	slice = rotate(slice, 5)
	fmt.Println("rotated slice is", slice)
}

func reverseArray(arr *[6]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func rotate(slice []int, positions int) []int {
	lastIdx := len(slice)
	firstIdx := lastIdx - (positions)
	subSlice1 := slice[firstIdx:lastIdx]
	subSlice2 := slice[0:firstIdx]

	return append(subSlice1, subSlice2...)
}
