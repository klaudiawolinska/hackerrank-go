package main

import (
	"fmt"
	"math"
)

/*
Given a square matrix, calculate the absolute difference between the sums of its diagonals.
*/

func diagonalDifference(arr [][]int32) int32 {
	d1 := int32(0)
	d2 := int32(0)

	for i := range arr {
		d1 += arr[i][i]
		d2 += arr[i][len(arr[0])-1-i]
	}

	return int32(math.Abs(float64(d1 - d2)))
}

func main() {
	// Expected output: 15
	m1 := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	fmt.Println(diagonalDifference(m1))
}
