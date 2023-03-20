package main

import (
	"fmt"
)

func main() {
	fmt.Println(TwoSum([]int{1, 2, 3}, 4))
}

func TwoSum(numbers []int, target int) [2]int {
	var res [2]int
	for i, v := range numbers {
		for j, v2 := range numbers {
			if v+v2 == target && i != j {
				res[0], res[1] = i, j
			}
		}
	}
	return res
}
