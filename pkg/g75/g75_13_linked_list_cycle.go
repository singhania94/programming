package main

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	t, h := head.Next, head.Next.Next

	for t != nil && h != nil && h.Next != nil {
		if t == h {
			return true
		}
		t = t.Next
		h = h.Next.Next
	}
	return false
}
