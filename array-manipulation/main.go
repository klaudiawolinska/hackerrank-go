// https://www.hackerrank.com/challenges/crush/problem

package main

import "fmt"

func main() {
	n := 40

	var matrix [][]int = [][]int{
		{19, 28, 419},
		{4, 23, 680},
		{5, 6, 907},
	}

	fmt.Println(arrayManipulation(n, matrix))
}

/*
// Not optimized for working with big input
func arrayManipulation(n int, arr [][]int) int {

	arrManipulated := make([]int, n)

	for _, manipulation := range arr {
		startIndex := manipulation[0]
		endIndex := manipulation[1]

		for i := startIndex - 1; i < endIndex; i++ {
			arrManipulated[i] += manipulation[2]
			//fmt.Println(arrManipulated)
		}
	}

	maxSum := arrManipulated[0]

	for i := range arrManipulated {
		if arrManipulated[i] > maxSum {
			maxSum = arrManipulated[i]
		}
	}

	return maxSum
}
*/

func arrayManipulation(n int, arr [][]int) int {

	arrManipulated := make([]int, n+1)

	for _, manipulation := range arr {
		startIndex := manipulation[0]
		endIndex := manipulation[1]

		arrManipulated[startIndex-1] += manipulation[2]
		arrManipulated[endIndex] -= manipulation[2]

		//fmt.Println(arrManipulated)
	}

	maxSum := 0
	currentSum := 0

	for i := range arrManipulated {
		currentSum += arrManipulated[i]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
