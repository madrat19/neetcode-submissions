/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	isSame := true
	dfs(p, q, &isSame)
	return isSame
}

func dfs(q, p *TreeNode, isSame *bool) {
	if q == nil && p == nil {
		return
	}
	if (q == nil && p != nil) || (q != nil && p == nil) {
		*isSame = false
		return
	}
	if q.Val != p.Val {
		*isSame = false
		return
	}
	dfs(p.Left, q.Left, isSame)
	dfs(p.Right, q.Right, isSame)
}
