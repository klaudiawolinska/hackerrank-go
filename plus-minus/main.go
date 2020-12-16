package main

import "fmt"

/*
Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero.
Print the decimal value of each fraction on a new line with 6 places after the decimal.
*/

func plusMinus(arr []int32) {
	positives := int32(0)
	negatives := int32(0)
	zeros := int32(0)

	for i := range arr {
		switch {
		case arr[i] > 0:
			positives++
		case arr[i] < 0:
			negatives++
		case arr[i] == 0:
			zeros++
		}
	}

	arrLength := int32(len(arr))

	fmt.Printf("%.6f\n", float32(positives)/float32(arrLength))
	fmt.Printf("%.6f\n", float32(negatives)/float32(arrLength))
	fmt.Printf("%.6f\n", float32(zeros)/float32(arrLength))
}

func main() {
	/* Expected output:
	   0.500000
	   0.333333
	   0.166667
	*/

	plusMinus([]int32{-4, 3, -9, 0, 4, 1})
}
