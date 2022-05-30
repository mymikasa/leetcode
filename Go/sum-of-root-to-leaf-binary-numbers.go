package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

//    0
//  1   2
func dfs(node *TreeNode, res int) int {
	if node == nil {
		return 0
	}
	res = res<<1 + node.Val

	if node.Left == nil && node.Right == nil {
		return res
	}

	return dfs(node.Left, res) + dfs(node.Right, res)

}
