func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lenght := 0
	node := head
	for node != nil {
		lenght++
		node = node.Next
	}
	number := lenght - n + 1

	dummy := &ListNode{Next: head}
	prev := dummy
	for i := 1; i < number; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
	return dummy.Next
}