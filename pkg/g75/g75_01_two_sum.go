package main

func twoSums(nums []int, target int) []int {
	// [target - val]index
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if j, ok := m[nums[i]]; ok {
			return []int{i, j}
		} else {
			m[target-nums[i]] = i
		}
	}
	return []int{}
}
