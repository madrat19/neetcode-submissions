/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	depth := 0
	depth = goDeep(root, depth)
	return depth
}

func goDeep(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	depth++
	depthLeft := goDeep(root.Left, depth)
	depthRight := goDeep(root.Right, depth)
	depth = max(depthLeft, depthRight)
	return depth
}
