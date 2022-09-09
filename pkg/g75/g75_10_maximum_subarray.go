package main

import "fmt"

func main() {
	fmt.Printf("maxSubArray: %v\n", maxSubArray([]int{5, 4, -1, 7, 8}))
}

func maxSubArray(nums []int) int {
	maxVal, curSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if curSum >= 0 {
			curSum = curSum + nums[i]
		} else {
			curSum = nums[i]
		}

		if curSum > maxVal {
			maxVal = curSum
		}
	}
	return maxVal
}
