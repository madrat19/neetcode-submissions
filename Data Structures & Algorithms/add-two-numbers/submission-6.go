func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry, val int
	prev := &ListNode{}
	dummy := prev
	for l1 != nil || l2 != nil {
		if l1 == nil {
			val = l2.Val + carry
		} else if l2 == nil {
			val = l1.Val + carry
		} else {
			val = l1.Val + l2.Val + carry
		}
		carry = val / 10
		val = val % 10
		node := &ListNode{Val: val}
		prev.Next = node
		prev = node
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		prev.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}