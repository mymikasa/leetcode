package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {

	var dfs func(root *TreeNode, level int, nodes map[int][]*TreeNode)
	nodes := make(map[int][]*TreeNode)

	dfs = func(root *TreeNode, level int, nodes map[int][]*TreeNode) {

		if level%2 == 1 {
			nodes[level] = append(nodes[level], root)
			// 0, 1, 2, 3, 4, 5
			ln := len(nodes[level])
			if ln == 1<<level {
				for i := 0; i < ln/2; i++ {
					temp := nodes[level][i].Val
					nodes[level][i].Val = nodes[level][ln-i-1].Val
					nodes[level][ln-i-1].Val = temp
				}
			}
		}
		if root.Left == nil {
			return
		}

		dfs(root.Left, level+1, nodes)
		dfs(root.Right, level+1, nodes)
	}
	dfs(root, 0, nodes)
	return root
}
