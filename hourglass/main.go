// https://www.hackerrank.com/challenges/2d-array/problem

package main

import "fmt"

func hourglassSum(arr [][]int) int {
	hourglassSums := []int{}

	for i, outer := range arr[:(len(arr) - 2)] {
		for j := range outer[:(len(arr) - 2)] {
			hourglass := arr[i][j] + arr[i][j+1] + arr[i][j+2]
			hourglass += arr[i+1][j+1]
			hourglass += arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			//fmt.Println(i, j)
			//fmt.Println(hourglassSum)
			hourglassSums = append(hourglassSums, hourglass)
		}
	}

	maxSum := hourglassSums[0]

	for i := range hourglassSums {
		if hourglassSums[i] > maxSum {
			maxSum = hourglassSums[i]
		}
	}

	return maxSum
}

func main() {
	var arr [][]int

	arr = [][]int{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}

	fmt.Println(hourglassSum(arr))
}
