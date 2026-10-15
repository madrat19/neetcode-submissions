/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := []*TreeNode{root}
	mp := map[*TreeNode]int{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		left := node.Left
		right := node.Right
		if left != nil {
			if _, ok := mp[left]; !ok {
				stack = append(stack, left)
				continue
			}
		}
		if right != nil {
			if _, ok := mp[right]; !ok {
				stack = append(stack, right)
				continue
			}
		}
		stack = stack[:len(stack)-1]

		leftHeight := mp[node.Left]
		rightHeight := mp[node.Right]
		if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
			return false
		}
		height := max(leftHeight, rightHeight) + 1
		mp[node] = height
	}
	return true
}
