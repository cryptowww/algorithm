package main

import (
	"cryptowww.leetcode/algo"
)

func main() {
	algo.RotateArrayToRightImplementByTempArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	algo.RotateArrayToRightImplementByReverse([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	algo.RotateArrayToLeftImplementByTempArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	algo.RotateArrayToLeftImplementByReverse([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
}
