package algo

import "strings"

// Given two strings s and goal, return true if and only if s can become goal after some number of shifts on s.
// A shift on s consists of moving the leftmost character of s to the rightmost position.
func RotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(s); i++ {
		if s[i] == goal[0] {
			j := 0
			for ; j < len(s); j++ {
				if s[(i+j)%len(s)] != goal[j] {
					break
				}
			}
			if j == len(s) {
				return true
			}
		}
	}
	return false
}

func RotateString2(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	return strings.Contains(s+s, goal)
}
