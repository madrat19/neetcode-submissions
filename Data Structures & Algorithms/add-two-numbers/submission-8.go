/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	var carry int
	for l1 != nil || l2 != nil || carry != 0 {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		value := v1 + v2 + carry
		carry = value / 10
		value = value % 10
		node := &ListNode{Val: value}
		prev.Next = node
		prev = node
	}
	return dummy.Next
}
