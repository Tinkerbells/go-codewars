// https://www.codewars.com/kata/54bf1c2cd5b56cc47f0007a1
package kata

import (
	"strings"
)

func duplicate_count(s1 string) int {
	s1 = strings.ToLower(s1)
	count := 0
	for _, char := range s1 {
		if strings.Index(s1, string(char)) != strings.LastIndex(s1, string(char)) {
			count++
			s1 = strings.ReplaceAll(s1, string(char), "")
		}
	}
	return count
}
