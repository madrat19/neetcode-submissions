/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	ans := sameTree(root, subRoot)
	if root != nil {
	ans = ans || (isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot))
	}
	return ans
}

func sameTree(root1, root2 *TreeNode) bool {
	ans := true
	if root1 == nil && root2 == nil {
		return true
	}
	if (root1 == nil && root2 != nil) || (root1 != nil && root2 == nil) {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	ans = sameTree(root1.Right, root2.Right) && sameTree(root1.Left, root2.Left)
	return ans 
}
