//given an array nums, rotate the array to the right by k steps, where k is non-negative

package algo

import "testing"

func TestRotateArrayToRightImplementByTempArray(t *testing.T) {
	result := RotateArrayToRightImplementByTempArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	if result[0] != 7 {
		t.Error("RotateArrayToRightImplementByTempArray failed, expected 7, got ", result[0])
	}
	result2 := RotateArrayToRightImplementByTempArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 13)
	if result2[0] != 6 {
		t.Error("RotateArrayToRightImplementByTempArray failed, expected 7, got ", result2[0])
	}
}

// rotate an array to the right implemented step by step, each step move one element to the right
func TestRotateArrayToRightImplementByReverse(t *testing.T) {
	result := RotateArrayToRightImplementByReverse([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)
	if result[0] != 7 {
		t.Error("RotateArrayToRightImplementByReverse failed, expected 7, got ", result[0])
	}
}
