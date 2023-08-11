package main

import (
	"fmt"

	"cryptowww.leetcode/algo"
)

func main() {
	fmt.Println(algo.RotateArrayToRightImplementByTempArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))
	fmt.Println(algo.RotateArrayToLeftImplementByTempArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))
}
