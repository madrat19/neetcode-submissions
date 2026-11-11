/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    ans := []int{}
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}

	for len(q) > 0 {
		qlen := len(q)
		ans = append(ans, q[qlen-1].Val)
		for i := 0; i < qlen; i++ {
			node := q[0]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			q = q[1:]
		}
	}
	return ans
}
