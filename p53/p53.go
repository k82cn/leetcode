package p53

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	preMax := nums[0]
	curMax := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if preMax+nums[i] > nums[i] {
			curMax = preMax + nums[i]
		} else {
			curMax = nums[i]
		}

		if curMax > max {
			max = curMax
		}

		preMax = curMax
	}

	return max
}
