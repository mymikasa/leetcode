package main

func findBottomLeftValue(root *TreeNode) int {
	cur := 0
	res := 0

	var dfs func(*TreeNode, int)

	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		height++
		dfs(node.Left, height)
		dfs(node.Right, height)

		if height > cur {
			res = node.Val
			cur = height
		}
	}
	dfs(root, 0)
	return res
}
