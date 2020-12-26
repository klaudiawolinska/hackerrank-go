// https://www.hackerrank.com/challenges/arrays-ds/problem

package main

import "fmt"

func reverseArray(a []int) []int {
	result := []int{}

	for i := range a {
		result = append(result, a[len(a)-i-1])
	}

	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(reverseArray(arr))
}
