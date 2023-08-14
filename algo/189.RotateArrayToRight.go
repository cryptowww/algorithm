//given an array nums, rotate the array to the right by k steps, where k is non-negative

package algo

// rotate an array to the right implement by temp array
func RotateArrayToRightImplementByTempArray(nums []int, k int) []int {
	n := len(nums)
	// if k is greater than n, then k = k % n
	if k > n {
		k = k % n
	}

	// save the last k elements to a temp array
	tempArr := make([]int, k)
	for i := n - k; i < n; i++ {
		tempArr[i-n+k] = nums[i]
	}

	// move the first n-k elements to the end
	//t := n - k - 1
	for i := n - k - 1; i >= 0; i-- {
		//nums[n-1-t+i] = nums[i]
		// at the fist cycle i = n-k-1, so k+i = n-1, while as i-- in every cycle, k+i will decrease 1
		nums[k+i] = nums[i]
	}

	// move the last k elements to the front
	for i := 0; i < k; i++ {
		nums[i] = tempArr[i]
	}

	return nums
}

// rotate an array to the right implemented step by step, each step move one element to the right
func RotateArrayToRightImplementByReverse(nums []int, k int) []int {
	n := len(nums)
	if k > n {
		k = k % n
	}
	for i := 0; i < k; i++ {
		temp := nums[n-1]

		for j := n - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = temp
	}
	return nums
}
