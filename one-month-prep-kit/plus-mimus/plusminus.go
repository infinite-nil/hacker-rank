package main

import "fmt"

func plusMinus(arr []int32) {
	// Save arr length for reuse
	size := len(arr)
	// Cast size type to float
	sizeFloat := float32(size)

	// First, let's create the variables that will hold the three kind of values and assign them 0.
	// Positives, negatives and zeros
	positive := 0
	negative := 0
	zero := 0

	// Loop through the array and verify each item
	for i := 0; i < size; i++ {
		if arr[i] > 0 {
			positive++
		} else if arr[i] < 0 {
			negative++
		} else {
			zero++
		}
	}

	// Print the result
	fmt.Println(float32(positive) / sizeFloat)
	fmt.Println(float32(negative) / sizeFloat)
	fmt.Println(float32(zero) / sizeFloat)
}

func main() {
	var arr []int32

	arr = append(arr, -4)
	arr = append(arr, 3)
	arr = append(arr, -9)
	arr = append(arr, 0)
	arr = append(arr, 4)
	arr = append(arr, 1)

	plusMinus(arr)
}
