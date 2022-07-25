package house_robber_three

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	memo := make(map[*TreeNode]int, 0)
	return solution(root, memo)
}

func solution(root *TreeNode, memo map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	val, ok := memo[root]
	if ok {
		return val
	}
	// 遍历根节点和下下个节点
	doSum := root.Val
	if root.Left != nil {
		doSum += solution(root.Left.Left, memo)
		doSum += solution(root.Left.Right, memo)
	}
	if root.Right != nil {
		doSum += solution(root.Right.Left, memo)
		doSum += solution(root.Right.Right, memo)
	}
	// 遍历下个节点
	notDoSum := 0
	notDoSum += solution(root.Left, memo)
	notDoSum += solution(root.Right, memo)
	res := max(doSum, notDoSum)
	memo[root] = res
	return res
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}
