func minOperations(nums []int, x int) int {
	numsSize := len(nums)
	target := getSliceSum(nums) - x
    
    fmt.Println(target)
	windowSize, curSum := -1, 0

	for left, right := 0, 0; right < numsSize; right++ {
		curSum += nums[right]
		for curSum > target && left <= right {
			curSum -= nums[left]
			left++
		}

		if curSum == target {
			windowSize = findMax(windowSize, right-left+1)
		}
	}

	if windowSize == -1 {
		return windowSize
	}

	return numsSize - windowSize
}

func getSliceSum(nums []int) int {
	curSum := 0
	for _, v := range nums {
		curSum += v
	}

	return curSum
}

func findMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

//read this

// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/discuss/1080880/Are-you-stuck-once-read-this-.-logic-explained-clearly-or-easy-to-understand

