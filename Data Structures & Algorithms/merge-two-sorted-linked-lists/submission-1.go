func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := list1
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list2.Val < list1.Val {
		head = list2
		list2 = list2.Next
	} else {
		list1 = list1.Next
	}

	node1 := list1
	node2 := list2
	prev := head
	for node1 != nil || node2 != nil {
		if node1 == nil {
			prev.Next = node2
			node2 = node2.Next
		} else if node2 == nil {
			prev.Next = node1
			node1 = node1.Next
		} else if node1.Val < node2.Val {
			prev.Next = node1
			node1 = node1.Next
		} else {
			prev.Next = node2
			node2 = node2.Next
		}
		prev = prev.Next
	}
	return head
}