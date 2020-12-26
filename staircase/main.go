// https://www.hackerrank.com/challenges/staircase/problem

package main

import (
	"fmt"
	"strings"
)

/*
Draw a staircase of size n.
Its base and height are both equal to n.
It is drawn using # symbols and spaces. The last line is not preceded by any spaces.
E.g. a staircase of size n = 4:
   #
  ##
 ###
####
*/

func staircase(n int32) {

	size := n

	for n > 0 {
		fmt.Print(strings.Repeat(" ", int(n-1)), strings.Repeat("#", int(size-n+1)), "\n")
		n--
	}
}

func main() {
	staircase(10)
}
