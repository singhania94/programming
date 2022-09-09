package main

func IsHappy(n int) bool {
	sums := make(map[int]bool, 0)
	return isHappyMap(n, sums)
}

func isHappyMap(n int, sums map[int]bool) bool {
	sum := 0
	for n != 0 {
		digit := n % 10
		sum = sum + digit*digit
		n = n / 10
	}
	if sum == 1 {
		return true
	} else if _, ok := sums[sum]; ok {
		return false
	}
	sums[sum] = true
	return isHappyMap(sum, sums)
}
