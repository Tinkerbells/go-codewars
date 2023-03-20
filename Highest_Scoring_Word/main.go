// https://www.codewars.com/kata/57eb8fcdf670e99d9b000272
package kata

import (
	"sort"
	"strings"
)

func High(s string) string {
	weights := []int{}
	strArray := strings.Split(s, " ")
	for _, word := range strings.Split(s, " ") {
		weight := 0
		for _, char := range string(word) {
			weight += int(char - 96)
		}
		weights = append(weights, weight)
	}
	sorted := make([]int, len(weights))
	copy(sorted, weights)
	sort.Ints(sorted)
	return strArray[indexOf(weights, sorted[len(sorted)-1])]
}

func indexOf(s []int, value int) int {
	for i, v := range s {
		if v == value {
			return i
		}
	}
	return -1
}
