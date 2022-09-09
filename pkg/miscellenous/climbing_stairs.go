package main

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}

	arr := make([]int, n)
	arr[0] = 1
	arr[1] = 2
	return climb(n, arr)
}

func climb(n int, arr []int) int {
	if arr[n-1] != 0 {
		return arr[n-1]
	}

	ans := climb(n-1, arr) + climb(n-2, arr)
	arr[n-1] = ans
	return ans
}
