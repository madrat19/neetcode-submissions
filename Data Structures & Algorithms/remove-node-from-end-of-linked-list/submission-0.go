/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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
	node = head
	for i := 1; i <= number; i++ {
		if i == number {
			prev.Next = node.Next
			return dummy.Next
		}
		prev = node
		node = node.Next
	}
	return dummy.Next
}
