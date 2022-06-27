package leetcode

func largestValues(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode, int)

	dfs = func(node *TreeNode, he int) {
		if node == nil {
			return
		}

		if he >= len(ans) {
			ans = append(ans, node.Val)
		} else {
			ans[he] = max(ans[he], node.Val)
		}

		dfs(node.Left, he+1)
		dfs(node.Right, he+1)
	}

	dfs(root, 0)
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
