package main

func BinarySearch(nums []int, target int) int {
	i, j := 0, len(nums)-1

	for i <= j {
		if nums[i] == target {
			return i
		} else if nums[j] == target {
			return j
		}

		mid := (i + j) / 2
		if nums[mid] <= target {
			if i == mid {
				return -1
			}
			i = mid
		} else {
			if i == mid {
				return -1
			}
			j = mid
		}
	}

	return -1
}
