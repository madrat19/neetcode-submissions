/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	good := 0
	var dfs func(node *TreeNode, mx int)
	dfs = func(node *TreeNode, mx int) {
		if node == nil {
			return
		}
		if node.Val >= mx {
			good++
		}
		mx = max(mx, node.Val)
		dfs(node.Left, mx)
		dfs(node.Right, mx)
	}
	dfs(root, root.Val)
	return good
}
