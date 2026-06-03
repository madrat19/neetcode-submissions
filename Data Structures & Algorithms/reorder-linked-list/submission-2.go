func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	list1 := head
	list2 := slow.Next
	slow.Next = nil
	list2 = reverseList(list2)

	for list2 != nil {
		temp1 := list1.Next
		temp2 := list2.Next
		list1.Next = list2
		list1 = temp1
		list2.Next = temp1
		list2 = temp2
	}
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}