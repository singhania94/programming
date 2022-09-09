package main

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	var ptr0 *ListNode
	ptr1 := head
	ptr2 := head
	for i := 0; i < n-1; i++ {
		ptr2 = ptr2.Next
	}
	for ptr2.Next != nil {
		ptr0 = ptr1
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	if ptr0 == nil {
		head = ptr1.Next
	} else {
		ptr0.Next = ptr1.Next
		ptr1 = nil
	}
	return head
}
