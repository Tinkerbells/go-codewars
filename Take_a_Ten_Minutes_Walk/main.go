// package kata

package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsValidWalk([]rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'}))
}

func IsValidWalk(walk []rune) bool {
	x, y := 0, 0
	if len(walk) != 10 {
		return false
	}
	for _, rune := range walk {
		switch rune {
		case 110:
			y += 1
		case 115:
			y -= 1
		case 119:
			x -= 1
		case 101:
			x += 1
		}
	}
	if x == 0 && y == 0 {
		return true
	}
	return false
}
