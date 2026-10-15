/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	balanced := true
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftHeight := dfs(node.Left)
		rightHeight := dfs(node.Right)
		if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
			balanced = false
		}
		return max(leftHeight, rightHeight) + 1
	}
	dfs(root)
	return balanced
}
