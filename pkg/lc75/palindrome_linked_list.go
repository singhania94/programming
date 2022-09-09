package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindrome(head *ListNode) bool {
	ptr := head
	arr := make([]int, 0)
	for ptr != nil {
		arr = append(arr, ptr.Val)
		ptr = ptr.Next
	}

	i, j := 0, len(arr)-1
	for i < j {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
