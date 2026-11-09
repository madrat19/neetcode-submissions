/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
    depth := 0
	dfs(depth, root, &ans)
	return ans
}

func dfs(depth int, node *TreeNode, ans *[][]int) {
	if node == nil {
		return 
	}
	depth++
	if len(*ans) < depth {
		arr := make([]int, 0, depth)
		*ans = append(*ans, arr)
	}
	(*ans)[depth-1] = append((*ans)[depth-1], node.Val)
	dfs(depth, node.Left, ans)
	dfs(depth, node.Right, ans)
}
