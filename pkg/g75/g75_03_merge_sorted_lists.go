package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	var ptr0 *ListNode
	ptr1 := list1
	ptr2 := list2

	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val <= ptr2.Val {
			ptr0 = ptr1
			ptr1 = ptr1.Next
			continue
		}
		next := ptr2.Next
		insert(ptr0, ptr1, ptr2)
		ptr0 = ptr2
		ptr2 = next
	}

	if ptr2 != nil {
		ptr0.Next = ptr2
	}

	if list1.Val <= list2.Val {
		return list1
	}
	return list2
}

func insert(ptr0, ptr1, ptr2 *ListNode) {
	if ptr0 != nil {
		ptr0.Next = ptr2
	}
	ptr2.Next = ptr1
}

// 3,9,11,12
// 4,5
