package unique_bst_two

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{nil}
	}
	return generate(1, n)
}


func generate(left, right int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if left > right {
		res = append(res, nil)
		return res
	}

	for i:= left; i<=right;i++ {
		leftSolutions := generate(left, i-1)
		rightSolutions := generate(i+1, right)

		for _, leftTree := range leftSolutions {
			for _, rightTree := range rightSolutions {
				tmp := &TreeNode{
					Val: i,
					Left: leftTree,
					Right: rightTree,
				}
				res = append(res, tmp)
			}
		}
	}
	return res
}
