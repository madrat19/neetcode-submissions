/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	hash := map[*Node]*Node{}
	oldNode := head
	for oldNode != nil {
		newNode := &Node{Val: oldNode.Val}
		hash[oldNode] = newNode
		oldNode = oldNode.Next
	}

	for newNode, oldNode := range hash {
		oldNode.Next = hash[newNode.Next]
		oldNode.Random = hash[newNode.Random]
	}
	return hash[head]
}
