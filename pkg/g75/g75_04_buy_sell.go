package main

func MaxProfit(prices []int) int {
	stack := make([]int, 0)
	var profit int

	stack = append(stack, prices[len(prices)-1])
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] >= stack[len(stack)-1] {
			stack = append(stack, prices[i])
		}
	}

	for i := 0; i < len(prices); i++ {
		if prices[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if profit < stack[len(stack)-1]-prices[i] {
			profit = stack[len(stack)-1] - prices[i]
		}
	}
	return profit
}
