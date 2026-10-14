/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	diameter := 0
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

		height := max(mp[node.Left], mp[node.Right]) + 1
		mp[node] = height
		diameter = max(diameter, mp[node.Left]+mp[node.Right])
	}
	return diameter
}
