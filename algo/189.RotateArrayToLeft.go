//given an array nums, rotate the array to the right by k steps, where k is non-negative

package algo

// rotate an array to the left implemented by temp array
func RotateArrayToLeftImplementByTempArray(nums []int, k int) []int {
	n := len(nums)

	if k > n {
		k = k % n
	}

	// save the first k elements
	temp := make([]int, k)
	for i := 0; i < k; i++ {
		temp[i] = nums[i]
	}

	// move the rest elements to the left
	for i := k; i < n; i++ {
		nums[i-k] = nums[i]
	}

	// move the first k elements to the end
	for i := 0; i < k; i++ {
		nums[n-k+i] = temp[i]
	}

	return nums
}

// rotate an array to the left implemented by reverse
func RotateArrayToLeftImplementByReverse(nums []int, k int) []int {
	n := len(nums)
	if k > n {
		k = k % n
	}

	// cycle k times
	for i := 0; i < k; i++ {
		temp := nums[0]
		// move the rest elements to the left
		for j := 1; j < n; j++ {
			nums[j-1] = nums[j]
		}

		// move the first element to the end
		nums[n-1] = temp
	}

	return nums
}
