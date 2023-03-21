// https://www.codewars.com/kata/5552101f47fc5178b1000050/go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(DigPow(46288, 3))
}

func DigPow(n, p int) int {
	a := splitInt(n)
	var sum float64 = 0.0
	for i, n := range a {
		sum += math.Pow(float64(n), float64(p+i))
	}
	if int(sum)%n == 0 {
		return int(sum) / n
	}
	return -1
}

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	for i, j := 0, len(slc)-1; i < j; i, j = i+1, j-1 {
		slc[i], slc[j] = slc[j], slc[i]
	}
	return slc
}
